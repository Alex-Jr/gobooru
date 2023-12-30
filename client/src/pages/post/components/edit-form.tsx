import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Select,
  Textarea,
  useColorModeValue,
  useToast,
} from "@chakra-ui/react";
import { useQueryClient } from "@tanstack/react-query";
import { Controller, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { useEditPost } from "services/posts/use-edit-post";
import { InputWithTag } from "shared/components/forms/input-with-tag";

interface IFormData {
  custom: string[];
  description: string;
  id: number;
  rating: string;
  sources: string[];
  tags: string;
  status: string;
}

interface IEditFormProps {
  isVisible: boolean;
  defaultValues: IFormData;
  onClose: () => void;
}
export const EditForm = ({
  isVisible,
  defaultValues,
  onClose,
}: IEditFormProps) => {
  const {
    register,
    control,
    formState: { errors },
    handleSubmit,
  } = useForm({
    defaultValues,
  });

  const editPostMutation = useEditPost();
  const queryClient = useQueryClient();
  const toast = useToast();

  const onSubmit = ({
    custom,
    description,
    id,
    rating,
    sources,
    status,
    tags,
  }: IFormData) => {
    editPostMutation.mutate(
      {
        custom,
        description,
        id,
        rating,
        sources,
        status,
        tags: tags.split(" "),
      },
      {
        onSuccess: (d) => {
          queryClient.invalidateQueries(["posts"]);
          onClose();
          toast({
            status: "success",
            description: t("feedback.editSuccess", {
              target: t("glossary.post"),
            }),
          });
        },
        onError: () => {
          toast({
            status: "error",
            description: t("feedback.editError", {
              target: t("glossary.post"),
            }),
          });
        },
      }
    );
  };

  const { t } = useTranslation();

  return (
    <Flex
      as="form"
      p={4}
      gap={4}
      mx={"auto"}
      w={"100%"}
      direction={"column"}
      hidden={!isVisible}
      bgColor={useColorModeValue("lightgray", "gray.700")}
      onSubmit={handleSubmit(onSubmit)}
    >
      <FormControl isInvalid={!!errors.tags}>
        <FormLabel htmlFor={"tags"}>{t("glossary.tags")}</FormLabel>

        {/* <Controller
          control={control}
          name={"tags"}
          render={({ field }) => <InputWithTag {...field} />}
        /> */}
        <Textarea {...register("tags")} h={200} />

        <FormErrorMessage>{errors.tags?.message}</FormErrorMessage>
      </FormControl>

      <FormControl isInvalid={!!errors.sources}>
        <FormLabel htmlFor={"sources"}>{t("glossary.sources")}</FormLabel>

        <Controller
          control={control}
          name={"sources"}
          render={({ field }) => <InputWithTag {...field} />}
        />

        <FormErrorMessage>{errors.tags?.message}</FormErrorMessage>
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

        <Textarea h={225} {...register("description")} />

        <FormErrorMessage>{errors.description?.message}</FormErrorMessage>
      </FormControl>

      <FormControl isInvalid={!!errors.custom}>
        <FormLabel htmlFor={"custom"}>{t("glossary.custom")}</FormLabel>

        <Controller
          control={control}
          name={"custom"}
          render={({ field }) => <InputWithTag {...field} />}
        />

        <FormErrorMessage>{errors.custom?.message}</FormErrorMessage>
      </FormControl>

      <Flex gap={2} mt={2}>
        <Button type="reset" colorScheme="yellow" flex={1}>
          {t("glossary.clear")}
        </Button>

        <Button type="submit" colorScheme="green" flex={1}>
          {t("glossary.submit")}
        </Button>
      </Flex>
    </Flex>
  );
};
