import { Box, Image } from "@chakra-ui/react";
import mime from "mime";

import { filePathToUrl } from "shared/functions/file_path_to_url";
import { imageModeToChakraCss } from "shared/functions/image-mode-to-chakra-css";
import { ImageModeEnum } from "shared/types/enums/image-mode-enum";

export const FilePreview = ({
  filePath,
  imageMode,
}: {
  filePath: string;
  imageMode: ImageModeEnum;
}) => {
  const mimeType = mime.getType(filePath);

  if (!mimeType) return <></>;

  if (mimeType.startsWith("video"))
    return (
      <Box
        as={"video"}
        controls
        loop
        src={filePathToUrl(filePath)}
        objectFit={"contain"}
      />
    );

  if (mimeType.startsWith("image"))
    return (
      <Image
        src={filePathToUrl(filePath)}
        objectFit="contain"
        {...imageModeToChakraCss(imageMode)}
        mx={"auto"}
      />
    );

  return <></>;
};
