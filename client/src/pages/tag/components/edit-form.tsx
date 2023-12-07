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
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { useTagCategoryList } from "services/tag-categories/use-tag-category";
import { useEditTag } from "services/tags/use-edit-tag";

interface IFormData {
  id: string;
  description: string;
  categoryId: string;
}

interface IEditFormProps {
  isOpen: boolean;
  defaultValues: IFormData;
  onClose: () => void;
}
export const EditForm = ({
  isOpen,
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

  const editTagMutation = useEditTag();
  const queryClient = useQueryClient();
  const toast = useToast();
  const { tagCategories } = useTagCategoryList();

  const onSubmit = (formData: IFormData) => {
    editTagMutation.mutate(formData, {
      onSuccess: (d) => {
        queryClient.invalidateQueries();
        toast({
          status: "success",
          description: t("feedback.editSuccess", {
            target: t("glossary.tag"),
          }),
        });
        onClose();
      },
      onError: () => {
        toast({
          status: "error",
          description: t("feedback.editError", { target: t("glossary.post") }),
        });
      },
    });
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
      hidden={!isOpen}
      bgColor={useColorModeValue("lightgray", "gray.700")}
      onSubmit={handleSubmit(onSubmit)}
    >
      <FormControl isInvalid={!!errors.categoryId}>
        <FormLabel htmlFor={"category"}>{t("glossary.category")}</FormLabel>

        <Select {...register("categoryId")}>
          {tagCategories.map((tc) => (
            <option key={`categoryId-o-${tc.id}`}>{tc.id}</option>
          ))}
        </Select>
      </FormControl>

      <FormControl isInvalid={!!errors.description}>
        <FormLabel htmlFor={"description"}>
          {t("glossary.description")}
        </FormLabel>

        <Textarea h={225} {...register("description")} />

        <FormErrorMessage>{errors.description?.message}</FormErrorMessage>
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
