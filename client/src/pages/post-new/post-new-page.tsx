import { AddIcon } from "@chakra-ui/icons";
import {
  Button,
  Checkbox,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Select,
  Textarea,
  useColorModeValue,
  useToast,
} from "@chakra-ui/react";
import { yupResolver } from "@hookform/resolvers/yup";
import { useEffect, useState } from "react";
import { Controller, useFieldArray, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";
import * as Yup from "yup";

import { useCreatePost } from "services/posts/use-create-post";
import { usePostMD5 } from "services/posts/use-post-md5";
import { hasNWords } from "shared/functions/yup/hasWords";
import { isFileType } from "shared/functions/yup/isFileType";
import { RatingEnum } from "shared/types/enums/rating-enum";

import { CustomInput } from "./components/custom-input";
import { ImageUploadInput } from "./components/image-upload-input";
import { SourcesInput } from "./components/sources-input";

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

const formSchema = Yup.object().shape({
  tags: Yup.string().required().test(hasNWords(1)),
  sources: Yup.array(
    Yup.object()
      .shape({
        source: Yup.string().required(),
      })
      .required()
  ).required(),
  file: Yup.mixed((i): i is File => i instanceof File)
    .required()
    .test(isFileType(validFiles)),
  description: Yup.string(),
  rating: Yup.string().oneOf(Object.values(RatingEnum)).required(),
  customs: Yup.array(
    Yup.object()
      .shape({
        custom: Yup.string().required(),
      })
      .required()
  ),
  parentId: Yup.string().matches(/^$|\d+$/),
  redirect: Yup.boolean().required(),
});

export const PostNewPage = () => {
  const [fileHash, setFileHash] = useState("");

  const post = usePostMD5({ md5: fileHash });

  const { t } = useTranslation();
  const createPostMutation = useCreatePost();

  const {
    control,
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm({
    mode: "onBlur",
    defaultValues: {
      tags: "",
      sources: [],
      rating: RatingEnum.SAFE,
      description: "",
      file: undefined,
      customs: [],
      parentId: "",
      redirect: true,
    },
    resolver: yupResolver(formSchema),
  });

  const {
    fields: sourceFields,
    append: sourceAppend,
    remove: sourceRemove,
  } = useFieldArray({
    control,
    name: "sources",
  });

  const {
    fields: customFields,
    append: customAppend,
    remove: customRemove,
  } = useFieldArray({
    control,
    name: "customs",
  });

  const toast = useToast();

  const navigate = useNavigate();

  useEffect(() => {
    if (post) {
      setError("file", { message: t("postNewPage.alreadyExist")! });
    }
  }, [post]);

  return (
    <Flex
      as={"form"}
      gap={4}
      overflow={"auto"}
      direction={{ sm: "column-reverse", md: "row" }}
      maxH={"90vh"}
      onSubmit={handleSubmit(async (data) => {
        if (post) return;

        const { post: createdPost } = await createPostMutation.mutateAsync({
          tags: data.tags.includes(",")
            ? data.tags
                .split(/,|\n/)
                .filter((w) => w.trim())
                .map((w) =>
                  w
                    .toLocaleLowerCase()
                    .split(" ")
                    .filter((w) => w.trim() !== "")
                    .join("_")
                )
            : data.tags
                .split(/\s+/)
                .filter((t) => t.trim() !== "")
                .map((w) => w.toLocaleLowerCase()),
          customs: data.customs?.map((c) => c.custom) || [],
          description: data.description || "",
          file: data.file,
          rating: data.rating,
          sources: data.sources.map((s) => s.source) || [],
          parentId: data.parentId,
        });

        toast({
          status: "success",
          description: t("feedback.creationSuccess", {
            target: t("glossary.post"),
          }),
        });

        if (data.redirect) {
          navigate(`/posts/${createdPost.id}`);
        }
      })}
    >
      <Flex
        direction={"column"}
        flex={0.5}
        gap={2}
        p={4}
        maxW={{ sm: "unset", lg: "550px" }}
        bgColor={useColorModeValue("lightgray", "gray.700")}
        h={"fit-content"}
      >
        <FormControl isInvalid={!!errors.tags}>
          <FormLabel htmlFor={"tags"}>{t("glossary.tags")}</FormLabel>

          <Textarea h={175} {...register("tags")} />

          <FormErrorMessage>{errors.tags?.message}</FormErrorMessage>
        </FormControl>

        <FormControl
          as={Flex}
          direction={"column"}
          isInvalid={!!errors.sources}
        >
          <FormLabel htmlFor={"sources"}>{t("glossary.sources")}</FormLabel>

          <Flex direction={"column"} gap={2}>
            {sourceFields.map((sf, i) => (
              <FormControl key={sf.id} isInvalid={!!errors.sources?.at!(i)}>
                <SourcesInput
                  index={i}
                  value={sf}
                  remove={sourceRemove}
                  register={() => register(`sources.${i}.source`)}
                />
                <FormErrorMessage alignItems={"start"}>
                  {errors.sources?.at!(i)?.source?.message}
                </FormErrorMessage>
              </FormControl>
            ))}

            <Button
              w={"fit-content"}
              rightIcon={<AddIcon />}
              onClick={() => sourceAppend({ source: "" })}
            >
              Add source
            </Button>
          </Flex>

          <FormErrorMessage>{errors.sources?.message}</FormErrorMessage>
        </FormControl>

        <FormControl isInvalid={!!errors.rating}>
          <FormLabel htmlFor={"rating"}>{t("glossary.rating")}</FormLabel>

          <Select {...register("rating")}>
            <option value="S">{t("glossary.safe")}</option>
            <option value="Q">{t("glossary.questionable")}</option>
            <option value="E">{t("glossary.explicit")}</option>
          </Select>

          <FormErrorMessage>{errors.rating?.message}</FormErrorMessage>
        </FormControl>

        <FormControl isInvalid={!!errors.description}>
          <FormLabel htmlFor={"description"}>
            {t("glossary.description")}
          </FormLabel>

          <Textarea h={175} {...register("description")} />

          <FormErrorMessage>{errors.description?.message}</FormErrorMessage>
        </FormControl>

        <FormControl as={Flex} direction={"column"}>
          <FormLabel htmlFor={"custom"}>{t("glossary.custom")}</FormLabel>

          <Flex direction={"column"} gap={2}>
            {customFields.map((cf, i) => (
              <FormControl key={cf.id} isInvalid={!!errors.customs?.at!(i)}>
                <CustomInput
                  index={i}
                  value={cf}
                  remove={customRemove}
                  register={() => register(`customs.${i}.custom`)}
                />
                <FormErrorMessage alignItems={"start"}>
                  {errors.customs?.at!(i)?.custom?.message}
                </FormErrorMessage>
              </FormControl>
            ))}

            <Button
              w={"fit-content"}
              rightIcon={<AddIcon />}
              onClick={() => customAppend({ custom: "" })}
            >
              Add custom
            </Button>
          </Flex>
        </FormControl>

        <FormControl isInvalid={!!errors.parentId}>
          <FormLabel htmlFor={"parentId"}>{t("glossary.parentId")}</FormLabel>

          <Input {...register("parentId")} />

          <FormErrorMessage>{errors.parentId?.message}</FormErrorMessage>
        </FormControl>

        <FormControl>
          <FormLabel htmlFor={"redirect"}>{t("glossary.redirect")}</FormLabel>

          <Checkbox {...register("redirect")} />
        </FormControl>

        <Flex gap={2} mt={2} position={"sticky"}>
          <Button type="reset" colorScheme="yellow" flex={1}>
            {t("glossary.clear")}
          </Button>

          <Button
            type="submit"
            colorScheme="green"
            flex={1}
            isDisabled={!!post}
          >
            {t("glossary.submit")}
          </Button>
        </Flex>
      </Flex>

      <FormControl isInvalid={!!post || !!errors.file} flex={1}>
        <Controller
          control={control}
          name={"file"}
          render={({ field: { ref, ...field } }) => (
            <ImageUploadInput {...field} onHash={setFileHash} />
          )}
        />

        <FormErrorMessage>{errors.file?.message}</FormErrorMessage>
      </FormControl>
    </Flex>
  );
};
