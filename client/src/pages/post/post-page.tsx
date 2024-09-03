import {
  Box,
  Flex,
  Text,
  Tooltip,
  useBoolean,
  useDisclosure,
} from "@chakra-ui/react";
import { useEffect, useRef, useState } from "react";
import { useParams } from "react-router-dom";

import { usePost } from "services/posts/use-post";
import { filePathToUrl } from "shared/functions/file_path_to_url";
import { ImageModeEnum } from "shared/types/enums/image-mode-enum";

import { DeleteModal } from "./components/delete-modal";
import { EditForm } from "./components/edit-form";
import { FilePreview } from "./components/file-preview";
import { PoolPagination } from "./components/pool-pagination";
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

  const [isNoteOpen, setNoteOpen] = useBoolean(true);

  const [imageMode, setImageMode] = useState<ImageModeEnum>(
    ImageModeEnum.FIT_V
  );

  const [isImageLoaded, setIsImageLoaded] = useBoolean(false);

  const [filePreviewDimensions, setFilePreviewDimensions] = useState({
    width: 0,
    height: 0,
  });

  const { id } = useParams() as { id: string };

  const post = usePost({ id });

  const filePreviewRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (filePreviewRef.current) {
      const { height, width } = filePreviewRef.current.getBoundingClientRect();
      setFilePreviewDimensions({ height, width });
    }
  }, [isImageLoaded, imageMode, post]);

  if (!post) return <></>;

  const heightScale = filePreviewDimensions.height / post.height;
  const widthScale = filePreviewDimensions.width / post.width;

  const onDownloadClick = () => {
    const link = document.createElement("a");
    link.href = filePathToUrl(post.file_path);
    link.download = post.md5;
    link.click();
  };

  return (
    <Flex direction={{ base: "column-reverse", lg: "row" }} gap={4}>
      <Sidebar
        post={post}
        imageMode={imageMode}
        setImageMode={setImageMode}
        onEditClick={onEditToggle}
        onDeleteClick={onDeleteToggle}
        onDownloadClick={onDownloadClick}
      />

      <Flex direction={"column"} gap={4} flex={1}>
        <PoolPagination post={post} />

        <Box
          position={"relative"}
          ref={filePreviewRef}
          w={"fit-content"}
          mx={"auto"}
          onClick={setNoteOpen.toggle}
          onLoad={() => {
            setIsImageLoaded.on();
          }}
        >
          <FilePreview filePath={post.file_path} imageMode={imageMode} />

          {isNoteOpen &&
            post.notes.map((note) => (
              <Tooltip key={note.id} label={note.body}>
                <Box
                  w={`${note.width * widthScale}px`}
                  h={`${note.height * heightScale}px`}
                  bgColor={"#ffe"}
                  opacity={0.5}
                  position={"absolute"}
                  top={note.y * heightScale}
                  left={note.x * widthScale}
                />
              </Tooltip>
            ))}
        </Box>

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
