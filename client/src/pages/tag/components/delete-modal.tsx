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

import { useDeleteTag } from "services/tags/use-delete-tag";
import { APITag } from "shared/types/services/tags/APITag";

export interface DeleteModalProps {
  tag: APITag["tag"];
  isOpen: boolean;
  onClose: () => void;
}

export const DeleteModal = ({
  tag: { id },
  isOpen,
  onClose,
}: DeleteModalProps) => {
  const deleteTagMutation = useDeleteTag();
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { t } = useTranslation();
  const toast = useToast();

  const handleDelete = () =>
    deleteTagMutation.mutate(
      { id },
      {
        onSuccess: () => {
          onClose();
          queryClient.invalidateQueries();
          toast({
            status: "success",
            description: t("feedback.deleteSuccess", {
              target: t("glossary.tag"),
            }),
          });
          navigate("/tags");
        },
        onError: () => {
          toast({
            status: "error",
            description: t("feedback.deleteError", {
              target: t("glossary.tag"),
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
            <Text color={"red.500"} fontWeight={"bold"}>
              {t("tagPage.deleteTag")}
            </Text>{" "}
            {t("tagPage.deleteTagMsg")}
          </Text>
        </ModalBody>

        <ModalFooter>
          <Button colorScheme="green" mr={3} onClick={onClose}>
            {t("glossary.cancel")}
          </Button>

          <Button colorScheme="red" mr={3} onClick={handleDelete}>
            {t("glossary.confirm")}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
