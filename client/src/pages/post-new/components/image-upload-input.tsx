import { CloseIcon } from "@chakra-ui/icons";
import { Flex, Icon, IconButton, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useDropzone } from "react-dropzone";
import { FaFileUpload } from "react-icons/fa";
import { ArrayBuffer } from "spark-md5";

import { FilePreview } from "./file-preview";

interface ImageUploadInputProps {
  onHash?: (hash: string) => void;
  onChange: (...event: any[]) => void;
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
  const [file, setFile] = useState<File | undefined>();

  const { getRootProps, isDragActive } = useDropzone({
    onDrop: (acceptedFiles: File[]) => {
      setFile(acceptedFiles[0]);
    },
    onError: console.error,
    accept: { "image/*": [".jpeg", ".png"], "video/*": [] },
  });

  useEffect(() => {
    onChange(file);
    onBlur();

    if (onHash) {
      if (file) {
        const blobSlice = File.prototype.slice,
          chunkSize = 2097152, // Read in chunks of 2MB
          chunks = Math.ceil(file.size / chunkSize),
          spark = new ArrayBuffer(),
          fileReader = new FileReader();

        let currentChunk = 0;

        fileReader.onload = (e: any) => {
          spark.append(e.target.result); // Append array buffer
          currentChunk++;

          if (currentChunk < chunks) {
            loadNext();
          } else {
            onHash(spark.end());
          }
        };

        fileReader.onerror = function () {
          console.warn();
        };

        const loadNext = () => {
          const start = currentChunk * chunkSize,
            end =
              start + chunkSize >= file.size ? file.size : start + chunkSize;

          fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
        };

        loadNext();
      } else {
        onHash("");
      }
    }
  }, [file]);

  if (!file)
    return (
      <Flex
        {...getRootProps()}
        direction="column"
        p={5}
        alignItems={"center"}
        justifyContent={"center"}
        border={"1px dashed gray"}
        flex={1}
        h={"95%"}
      >
        <Icon as={FaFileUpload} w="30px" h="40px" mb={5}></Icon>

        <Text fontSize={"sm"}>
          {isDragActive ? "Drag active" : "Drag inactive"}
        </Text>
      </Flex>
    );

  return (
    <Flex
      direction={"column"}
      gap={2}
      flex={1}
      overflow={{ sm: "none", md: "auto" }}
      w={"fit-content"}
      maxW={"70vw"}
      maxH={"90vh"}
    >
      <Flex alignItems={"center"} justify={"space-between"}>
        <Text>{file.name}</Text>

        <IconButton
          aria-label="Clear"
          icon={<CloseIcon />}
          onClick={() => setFile(undefined)}
          size={"sm"}
        />
      </Flex>

      <FilePreview file={file} />
    </Flex>
  );
};
