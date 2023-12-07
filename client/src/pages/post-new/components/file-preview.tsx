import { Box, Image } from "@chakra-ui/react";

export interface FilePreviewProps {
  file: File;
}

export const FilePreview = ({ file }: FilePreviewProps) => {
  if (file.type.startsWith("video"))
    return (
      <Box
        as={"video"}
        controls
        src={URL.createObjectURL(file)}
        objectFit={"contain"}
      />
    );

  if (file.type.startsWith("image"))
    return (
      <Image
        src={URL.createObjectURL(file)}
        alt="Preview"
        objectFit={"contain"}
      />
    );

  return <Box>File type not supported</Box>;
};
