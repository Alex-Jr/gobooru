import { Flex, Text, useDisclosure } from "@chakra-ui/react";
import { useState } from "react";
import { useParams } from "react-router-dom";

import { usePost } from "services/posts/use-post";
import { ImageModeEnum } from "shared/types/enums/image-mode-enum";

import { DeleteModal } from "./components/delete-modal";
import { EditForm } from "./components/edit-form";
import { FilePreview } from "./components/file-preview";
import { PoolPagination } from "./components/pool-pagination";
import { PostRelations } from "./components/post-relations";
import { Sidebar } from "./components/sidebar";

export const PostPage = () => {
  const {
    isOpen: isEditOpen,
    onClose: onEditClose,
    onToggle: onEditToggle,
  } = useDisclosure();

  const {
    isOpen: isDeleteOpen,
    onOpen: onDeleteOpen,
    onClose: onDeleteClose,
    onToggle: onDeleteToggle,
  } = useDisclosure();

  const [imageMode, setImageMode] = useState<ImageModeEnum>(
    ImageModeEnum.FIT_H
  );

  const { id } = useParams() as { id: string };

  const post = usePost({ id });

  if (!post) return <></>;

  return (
    <Flex direction={{ base: "column-reverse", lg: "row" }} gap={4}>
      <Sidebar
        post={post}
        imageMode={imageMode}
        setImageMode={setImageMode}
        onEditClick={onEditToggle}
        onDeleteClick={onDeleteToggle}
      />

      <Flex direction={"column"} gap={4} flex={1}>
        <PoolPagination post={post} />

        <FilePreview filePath={post.file_path} imageMode={imageMode} />

        <PostRelations relations={post.relations} />

        {post.description && (
          <Flex bgColor={"gray.700"} p={4} direction={"column"}>
            <Text fontSize={"2xl"} mb={2}>
              Description:
            </Text>
            {post.description}
          </Flex>
        )}

        <EditForm
          isVisible={isEditOpen}
          defaultValues={{
            custom: post.custom,
            description: post.description,
            id: post.id,
            rating: post.rating,
            sources: post.sources,
            status: post.status as unknown as string,
            tags: post.tags.map((t) => t.id).join(" "),
          }}
          onClose={onEditClose}
        />

        <DeleteModal
          post={post}
          isOpen={isDeleteOpen}
          onClose={onDeleteClose}
        />
      </Flex>
    </Flex>
  );
};
