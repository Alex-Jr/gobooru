import {
  Box,
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Input,
} from "@chakra-ui/react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

interface IFormInput {
  search: string;
}

export const SearchForm = () => {
  const { t } = useTranslation();

  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm<IFormInput>();

  const onSubmit: SubmitHandler<IFormInput> = (formData) => {
    console.log(formData);
  };

  return (
    <Box as={"form"} onSubmit={handleSubmit(onSubmit)}>
      <FormControl isInvalid={!!errors.search}>
        <Input
          {...register("search")}
          placeholder={t("forms.tagPlaceholder")!}
          variant={"Filled"}
          autoComplete="off"
        />
        <FormErrorMessage>{errors.search?.message}</FormErrorMessage>
      </FormControl>

      <Flex justify={"space-between"} mt={4} gap={4}>
        <Button flex={1} colorScheme="orange">
          {t("glossary.surpriseMe")}
        </Button>

        <Button type="submit" flex={1} colorScheme="orange">
          {t("glossary.search")}
        </Button>
      </Flex>
    </Box>
  );
};
