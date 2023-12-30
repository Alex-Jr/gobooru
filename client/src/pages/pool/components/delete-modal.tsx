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

import { useDeletePool } from "services/pools/use-delete-pool";

interface IDeleteModalProps {
  pool: { id: number };
  isOpen: boolean;
  onClose: () => void;
}
export const DeleteModal = ({
  pool: { id },
  isOpen,
  onClose,
}: IDeleteModalProps) => {
  const { t } = useTranslation();
  const deletePoolMutation = useDeletePool();
  const queryClient = useQueryClient();
  const toast = useToast();
  const navigate = useNavigate();

  const onDelete = () =>
    deletePoolMutation.mutate(
      { id },
      {
        onSuccess: (d) => {
          queryClient.invalidateQueries(["pool"]);
          onClose();
          toast({
            status: "success",
            description: t("feedback.deleteSuccess", {
              target: t("glossary.pool"),
            }),
          });
          navigate("/pools");
        },
        onError: () => {
          toast({
            status: "error",
            description: t("feedback.deleteError", {
              target: t("glossary.pool"),
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
              {t("glossary.deletePool")}
            </Text>
          </Text>
        </ModalBody>

        <ModalFooter>
          <Button colorScheme="green" mr={3} onClick={onClose}>
            {t("glossary.cancel")}
          </Button>

          <Button colorScheme="red" mr={3} onClick={onDelete}>
            {t("glossary.confirm")}
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
