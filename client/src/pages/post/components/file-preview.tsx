import { Box, Image } from "@chakra-ui/react";
import mime from "mime";
import { forwardRef } from "react";

import { filePathToUrl } from "shared/functions/file_path_to_url";
import { imageModeToChakraCss } from "shared/functions/image-mode-to-chakra-css";
import { ImageModeEnum } from "shared/types/enums/image-mode-enum";

interface IFilePreviewProps {
  filePath: string;
  imageMode: ImageModeEnum;
}

export const FilePreview = forwardRef<any, IFilePreviewProps>(
  function FilePreview({ filePath, imageMode }, ref) {
    const mimeType = mime.getType(filePath);

    if (!mimeType) return <></>;

    if (mimeType.startsWith("video"))
      return (
        <Box
          ref={ref}
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
          ref={ref}
          src={filePathToUrl(filePath)}
          objectFit="contain"
          {...imageModeToChakraCss(imageMode)}
        />
      );

    return <></>;
  }
);
