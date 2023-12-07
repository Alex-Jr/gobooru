import { Flex, Icon, Text } from "@chakra-ui/react";
import { useDropzone } from "react-dropzone";
import { FaFileUpload } from "react-icons/fa";

interface ImageUploadInputProps {
  onHash?: (hash: string) => void;
  onChange: (event: File[]) => void;
  onBlur: () => void;
  disabled?: boolean;
  value: File | undefined;
  name: string;
}

export const ImageUploadInput = ({
  name,
  onBlur,
  onChange,
  onHash,
}: ImageUploadInputProps) => {
  const { getRootProps, isDragActive } = useDropzone({
    onDrop: (acceptedFiles: File[]) => {
      onChange(acceptedFiles);
      onBlur();
    },
    onError: console.error,
    accept: { "image/*": [".jpeg", ".png"], "video/*": [] },
  });

  return (
    <Flex
      {...getRootProps()}
      direction="column"
      p={5}
      alignItems={"center"}
      justifyContent={"center"}
      border={"1px dashed gray"}
      flex={1}
      h={300}
    >
      <Icon as={FaFileUpload} w="30px" h="40px" mb={5}></Icon>

      <Text fontSize={"sm"}>
        {isDragActive ? "Drag active" : "Drag inactive"}
      </Text>
    </Flex>
  );
};
