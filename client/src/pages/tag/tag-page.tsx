import { DeleteIcon, EditIcon, StarIcon } from "@chakra-ui/icons";
import {
  Button,
  Divider,
  Flex,
  HStack,
  Text,
  useDisclosure,
} from "@chakra-ui/react";
import { useTranslation } from "react-i18next";
import { useParams } from "react-router-dom";

import { useTagCategoryList } from "services/tag-categories/use-tag-category";
import { useTag } from "services/tags/use-tag";

import { DeleteModal } from "./components/delete-modal";
import { EditForm } from "./components/edit-form";

export const TagPage = () => {
  const { id } = useParams();
  const { t } = useTranslation();
  const tag = useTag({ id: id! });
  const { tagCategoriesObj } = useTagCategoryList();

  const {
    isOpen: isDeleteOpen,
    onClose: onDeleteClose,
    onOpen: onDeleteOpen,
    onToggle: onDeleteToggle,
  } = useDisclosure();

  const {
    isOpen: isEditOpen,
    onClose: onEditClose,
    onOpen: onEditOpen,
    onToggle: onEditToggle,
  } = useDisclosure();

  if (!tag) return <></>;

  return (
    <Flex direction={"column"} bgColor={"gray.700"} p={4}>
      <Text fontSize={"4xl"}>{id}</Text>

      <Text fontSize="xl" color={tagCategoriesObj[tag.categoryId]}>
        {tag.categoryId}
      </Text>

      <Divider my={2} />

      {tag.description && (
        <>
          <Text fontSize="xl"> {tag.description} </Text>

          <Divider my={2} />
        </>
      )}

      <HStack w={"100%"} justifyContent={"flex-end"}>
        <Button rightIcon={<StarIcon />}>{t("glossary.favorite")}</Button>

        <Button rightIcon={<EditIcon />} onClick={onEditToggle}>
          {t("glossary.edit")}
        </Button>

        <Button rightIcon={<DeleteIcon />} onClick={onDeleteToggle}>
          {t("glossary.delete")}
        </Button>
      </HStack>

      <EditForm defaultValues={tag} isOpen={isEditOpen} onClose={onEditClose} />

      <DeleteModal tag={tag} isOpen={isDeleteOpen} onClose={onDeleteClose} />
    </Flex>
  );
};
