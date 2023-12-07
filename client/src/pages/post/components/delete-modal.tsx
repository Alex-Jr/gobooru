import {
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  useToast,
} from "@chakra-ui/react";
import { useQueryClient } from "@tanstack/react-query";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { useDeletePost } from "services/posts/use-delete-post";

interface IDeleteModalProps {
  post: { id: number };
  isOpen: boolean;
  onClose: () => void;
}
export const DeleteModal = ({
  post: { id },
  isOpen,
  onClose,
}: IDeleteModalProps) => {
  const { t } = useTranslation();
  const deletePostMutation = useDeletePost();
  const queryClient = useQueryClient();
  const toast = useToast();
  const navigate = useNavigate();

  const onDelete = (type: "HARD" | "SOFT") =>
    deletePostMutation.mutate(
      { id, type },
      {
        onSuccess: (d) => {
          queryClient.invalidateQueries(["posts"]);
          onClose();
          toast({
            status: "success",
            description: t("feedback.deleteSuccess", {
              target: t("glossary.post"),
            }),
          });
          navigate("/posts");
        },
        onError: () => {
          toast({
            status: "error",
            description: t("feedback.deleteError", {
              target: t("glossary.post"),
            }),
          });
        },
      }
    );

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />

      <ModalContent>
        <ModalHeader>{t("glossary.areYouSure")}</ModalHeader>

        <ModalCloseButton />

        <ModalBody>
          <Text>
            <Text color={"lightblue"} fontWeight={"bold"}>
              {t("glossary.deleteFile")}
            </Text>

            {t("postPage.softDelete")}
          </Text>
          <Text>
            <Text color={"red.500"} fontWeight={"bold"}>
              {t("glossary.deletePost")}
            </Text>

            {t("postPage.hardDelete")}
          </Text>
        </ModalBody>

        <ModalFooter>
          <Button colorScheme="green" mr={3} onClick={onClose}>
            {t("glossary.cancel")}
          </Button>

          <Button colorScheme="cyan" mr={3} onClick={() => onDelete("SOFT")}>
            {t("glossary.deleteFile")}
          </Button>

          <Button colorScheme="red" mr={3} onClick={() => onDelete("HARD")}>
            {t("glossary.deletePost")}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
