import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Textarea,
  useToast,
} from "@chakra-ui/react";
import { useQueryClient } from "@tanstack/react-query";
import { Controller, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { useEditPool } from "services/pools/use-edit-pool";
import { InputWithTag } from "shared/components/forms/input-with-tag";

interface IFormData {
  id: number;
  custom: string[];
  description: string;
  name: string;
  posts: string;
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

  const editPoolMutation = useEditPool();
  const queryClient = useQueryClient();
  const toast = useToast();

  const onSubmit = ({ id, custom, description, name, posts }: IFormData) => {
    editPoolMutation.mutate(
      {
        id,
        custom,
        description,
        name,
        posts: posts
          .trim()
          .split(" ")
          .map((p) => parseInt(p, 10)),
      },
      {
        onSuccess: (d) => {
          queryClient.invalidateQueries(["pools"]);
          onClose();
          toast({
            status: "success",
            description: t("feedback.editSuccess", {
              target: t("glossary.pool"),
            }),
          });
        },
        onError: () => {
          toast({
            status: "error",
            description: t("feedback.editError", {
              target: t("glossary.pool"),
            }),
          });
        },
      }
    );
  };

  const { t } = useTranslation();

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />

      <ModalContent>
        <ModalHeader>{t("glossary.areYouSure")}</ModalHeader>

        <ModalCloseButton />

        <ModalBody as={"form"} onSubmit={handleSubmit(onSubmit)}>
          <FormControl isInvalid={!!errors.name}>
            <FormLabel htmlFor={"name"}>{t("glossary.name")}</FormLabel>

            <Input {...register("name")} />

            <FormErrorMessage>{errors.name?.message}</FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={!!errors.posts}>
            <FormLabel htmlFor={"posts"}>{t("glossary.posts")}</FormLabel>

            <Textarea {...register("posts")} h={200} />

            <FormErrorMessage>{errors.posts?.message}</FormErrorMessage>
          </FormControl>

          <FormControl isInvalid={!!errors.description}>
            <FormLabel htmlFor={"description"}>
              {t("glossary.description")}
            </FormLabel>

            <Textarea {...register("description")} h={200} />

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
        </ModalBody>
        <ModalFooter>
          <Button colorScheme="green" mr={3} onClick={onClose}>
            {t("glossary.cancel")}
          </Button>

          <Button colorScheme="yellow" mr={3} onClick={handleSubmit(onSubmit)}>
            {t("glossary.confirm")}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
