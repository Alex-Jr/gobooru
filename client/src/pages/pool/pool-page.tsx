import { DeleteIcon, EditIcon, LinkIcon, StarIcon } from "@chakra-ui/icons";
import {
  Button,
  Divider,
  Flex,
  HStack,
  Tag,
  Text,
  VStack,
  useDisclosure,
} from "@chakra-ui/react";
import { useTranslation } from "react-i18next";
import { useParams } from "react-router-dom";

import { usePool } from "services/pools/use-pool";
import { GenericGrid } from "shared/components/lists/generic-list";
import { PostItem } from "shared/components/lists/post-item";

import { DeleteModal } from "./components/delete-modal";
import { EditForm } from "./components/edit-form";

export const PoolPage = () => {
  const { id } = useParams();
  const { t } = useTranslation();
  const pool = usePool({ id: id! });

  const {
    isOpen: isEditOpen,
    onToggle: onEditToggle,
    onClose: onEditClose,
  } = useDisclosure();
  const {
    isOpen: isDeleteOpen,
    onToggle: onDeleteToggle,
    onClose: onDeleteClose,
  } = useDisclosure();

  if (!pool) return <></>;

  return (
    <Flex direction={"column"} gap={4}>
      <VStack align={"start"} bgColor={"gray.700"} p={4}>
        <Text fontSize={"2xl"} fontWeight={"bold"}>
          {pool.name}
        </Text>

        <Text fontSize={"xl"}>Description: </Text>
        <Text textAlign={"justify"}>{pool.description}</Text>

        <Divider my={4} />

        <HStack>
          <Text fontSize={"xl"}>Post count:</Text>
          <Text>{pool.post_count}</Text>
        </HStack>

        <HStack>
          <Text fontSize={"xl"}>Creation date: </Text>
          <Text>{new Date(pool.created_at).toLocaleString()}</Text>
        </HStack>

        <HStack>
          <Text fontSize={"xl"}>Last Update: </Text>
          <Text>{new Date(pool.updated_at).toLocaleString()}</Text>
        </HStack>

        {pool.custom.length > 0 && <Text fontSize={"xl"}>Custom: </Text>}
        <HStack>
          {pool.custom.map((custom, index) => (
            <Tag key={`custom-${index}`}>{custom}</Tag>
          ))}
        </HStack>

        <Divider my={4} />

        <HStack w={"100%"} justifyContent={"space-between"}>
          <HStack>
            <Button rightIcon={<LinkIcon />}>{t("glossary.post")}</Button>
          </HStack>

          <HStack>
            <Button rightIcon={<StarIcon />} colorScheme="green">
              {t("glossary.favorite")}
            </Button>
            <Button
              rightIcon={<EditIcon />}
              onClick={onEditToggle}
              colorScheme="yellow"
            >
              {t("glossary.edit")}
            </Button>
            <Button
              rightIcon={<DeleteIcon />}
              onClick={onDeleteToggle}
              colorScheme="red"
            >
              {t("glossary.delete")}
            </Button>
          </HStack>
        </HStack>
      </VStack>

      <GenericGrid
        w="250px"
        items={pool.posts}
        renderItem={(post, index) => (
          <PostItem key={`pool-item-${index}`} post={post} />
        )}
      />

      <EditForm
        defaultValues={{
          id: pool.id,
          custom: pool.custom,
          description: pool.description,
          name: pool.name,
          posts: pool.posts.map((p) => p.id).join(" "),
        }}
        isOpen={isEditOpen}
        onClose={onEditClose}
      />
      <DeleteModal isOpen={isDeleteOpen} onClose={onDeleteClose} pool={pool} />
    </Flex>
  );
};
