import { ChevronDownIcon, ChevronUpIcon, DeleteIcon } from "@chakra-ui/icons";
import {
  Divider,
  Flex,
  FormControl,
  FormLabel,
  IconButton,
  Input,
  Select,
  Text,
  Textarea,
} from "@chakra-ui/react";

import { RatingEnum } from "shared/types/enums/rating-enum";

import { FilePreview } from "./file-preview";

interface IFormData {
  id: string;
  file: File;
  rating: RatingEnum;
  tags: string;
  description: string;
  sources: string;
}

interface PostFormProps {
  index: number;
  value: IFormData;
  remove: any;
  move: any;
  length: number;
  register: any;
}

export const PostForm = ({
  index,
  value,
  remove,
  move,
  length,
  register,
}: PostFormProps) => {
  return (
    <>
      <Flex gap={2} w={"100%"}>
        <Text>{index}°</Text>
        <Flex w={200} minW={200} h={200} justifyContent={"center"}>
          <FilePreview file={value.file} />
        </Flex>

        <FormControl
          as={Flex}
          flexDir={"column"}
          justifyContent={"space-between"}
        >
          <FormLabel htmlFor="tags">Tags</FormLabel>
          <Textarea {...register(`posts.${index}.tags`)} h={180} />
        </FormControl>

        <FormControl
          as={Flex}
          flexDir={"column"}
          justifyContent={"space-between"}
        >
          <FormLabel htmlFor="description">description</FormLabel>

          <Textarea {...register(`posts.${index}.description`)} h={180} />
        </FormControl>

        <Flex direction={"column"} minW={250} justifyContent={"space-between"}>
          <FormControl>
            <FormLabel htmlFor="rating">Rating</FormLabel>

            <Select {...register(`posts.${index}.rating`)}>
              {Object.entries(RatingEnum).map(([n, v]) => (
                <option value={v} key={`${value.id}-rating-${n}`}>
                  {n}
                </option>
              ))}
            </Select>
          </FormControl>

          <FormControl>
            <FormLabel htmlFor="sources">Source</FormLabel>

            <Input {...register(`posts.${index}.sources`)} />
          </FormControl>
        </Flex>

        <Flex direction={"column"} justifyContent={"space-between"} ml={4}>
          <IconButton
            aria-label="move-up"
            colorScheme="green"
            icon={<ChevronUpIcon />}
            size={"sm"}
            onClick={() => {
              if (index > 0) {
                move(index, index - 1);
              }
            }}
          />
          <IconButton
            aria-label="remove"
            icon={<DeleteIcon />}
            colorScheme="red"
            size={"sm"}
            onClick={() => remove(index)}
          />
          <IconButton
            aria-label="move-down"
            icon={<ChevronDownIcon />}
            colorScheme="green"
            size={"sm"}
            onClick={() => {
              if (index < length) {
                move(index, index + 1);
              }
            }}
          />
        </Flex>
      </Flex>
      <Divider />
    </>
  );
};
