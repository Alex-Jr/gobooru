import { Select } from "@chakra-ui/react";

import { ImageModeEnum } from "shared/types/enums/image-mode-enum";

interface IImageMode {
  value: ImageModeEnum;
  onChange: (x: ImageModeEnum) => void;
}

export const ImageModeSelect = ({ value, onChange }: IImageMode) => {
  return (
    <Select
      size={"sm"}
      value={value}
      onChange={(e) => {
        onChange(e.target.value as ImageModeEnum);
      }}
    >
      {Object.entries(ImageModeEnum).map(([, k]) => (
        <option key={`image-mode-${k}`} value={k}>
          {k}
        </option>
      ))}
    </Select>
  );
};
