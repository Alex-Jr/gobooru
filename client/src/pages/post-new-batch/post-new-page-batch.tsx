import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Input,
  Textarea,
  VStack,
} from "@chakra-ui/react";
import { useState } from "react";
import { useFieldArray, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { ImageUploadInput } from "pages/post-new-batch/components/image-upload-input";
import { useCreatePool } from "services/pools/use-create-pool";
import { useCreatePost } from "services/posts/use-create-post";
import { usePost } from "services/posts/use-post";
import { RatingEnum } from "shared/types/enums/rating-enum";

import { PostForm } from "./components/post-form";

const validFiles = [
  "image/png",
  "image/jpeg",
  "image/jpg",
  "image/gif",
  "image/webp",
  "video/webm",
  "video/quicktime",
  "video/mp4",
];

interface IFormData {
  posts: {
    file: File;
    rating: RatingEnum;
    tags: string;
    description: string;
    sources: string;
  }[];
  pool: string;
  description: string;
}

export const PostNewPageBatch = () => {
  const [fileHash, setFileHash] = useState("");

  const post = usePost({ id: fileHash });

  const { t } = useTranslation();

  const createPostMutation = useCreatePost();
  const createPoolMutation = useCreatePool();

  // const formSchema = Yup.object().shape({
  //   tags: Yup.array(Yup.string().required()).required().min(5),
  //   sources: Yup.array(Yup.string().required()).min(1).required(),
  //   file: Yup.array(
  //     Yup.mixed((i): i is File => i instanceof File)
  //       .required()
  //       .test(isFileType(validFiles))
  //   ),
  //   description: Yup.string().required().min(0),
  //   rating: Yup.string().oneOf(Object.values(RatingEnum)).required(),
  // });

  const {
    control,
    register,
    handleSubmit,
    setError,
    watch,
    getValues,
    formState: { errors },
  } = useForm<IFormData>({
    // mode: "onBlur",
    // resolver: yupResolver(formSchema),
  });

  const { fields, append, update, remove, move } = useFieldArray({
    control,
    name: "posts",
  });

  const navigate = useNavigate();

  return (
    <Flex
      as={"form"}
      gap={4}
      direction={"column"}
      onSubmit={handleSubmit(async (formData) => {
        const postsIds: number[] = [];

        for (const post of formData.posts) {
          const { post: createdPost } = await createPostMutation.mutateAsync({
            description: post.description,
            file: post.file,
            rating: post.rating,
            sources: post.sources.split(" "),
            tags: post.tags.includes(",")
              ? post.tags
                  .split(/,|\n/)
                  .filter((w) => w.trim())
                  .map((w) =>
                    w
                      .toLocaleLowerCase()
                      .split(" ")
                      .filter((w) => w.trim() !== "")
                      .join("_")
                  )
              : post.tags
                  .split(/\s+/)
                  .filter((t) => t.trim() !== "")
                  .map((w) => w.toLocaleLowerCase()),
            customs: [],
          });

          postsIds.push(createdPost.id);
        }
        if (formData.pool) {
          const { pool } = await createPoolMutation.mutateAsync({
            custom: [],
            name: formData.pool,
            description: formData.description,
            posts: postsIds,
          });

          navigate(`/pools/${pool.id}`);
        }
      })}
    >
      <Flex gap={8}>
        <Flex direction={"column"} flex={1} justifyContent={"space-between"}>
          <FormControl>
            <FormLabel htmlFor="pool">{t("glossary.pool")}</FormLabel>

            <Input {...register("pool")} />
          </FormControl>

          <FormControl>
            <FormLabel htmlFor="description">
              {t("glossary.description")}
            </FormLabel>

            <Textarea {...register("description")} />
          </FormControl>

          <Flex gap={4} placeSelf={"end"}>
            <Button
              mt={10}
              colorScheme="blue"
              w={200}
              alignSelf={"end"}
              onClick={() => {
                fields.forEach((f, i) => {
                  update(i, {
                    ...f,
                    rating: getValues("posts.0.rating"),
                    tags: getValues("posts.0.tags"),
                    sources: getValues("posts.0.sources"),
                  });
                });
              }}
            >
              {t("glossary.copy")}
            </Button>

            <Button
              mt={10}
              type={"reset"}
              colorScheme="yellow"
              w={200}
              alignSelf={"end"}
            >
              {t("glossary.clear")}
            </Button>

            <Button
              mt={10}
              type={"submit"}
              colorScheme="green"
              w={200}
              alignSelf={"end"}
            >
              {t("glossary.submit")}
            </Button>
          </Flex>
        </Flex>

        <ImageUploadInput
          value={undefined}
          name=""
          onBlur={() => undefined}
          onChange={(e) => {
            e.forEach((f) => {
              append({
                file: f,
                rating: RatingEnum.SAFE,
                description: "",
                tags: "",
                sources: "",
              });
            });
          }}
          onHash={setFileHash}
        />
      </Flex>

      <VStack gap={10} maxH={"90vh"} overflow={"auto"}>
        {fields.map((f, i) => (
          <PostForm
            key={f.id}
            index={i}
            value={f}
            remove={remove}
            move={move}
            length={fields.length}
            register={register}
          />
        ))}
      </VStack>
    </Flex>
  );
};
