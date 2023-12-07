import { DeleteIcon } from "@chakra-ui/icons";
import {
  Flex,
  IconButton,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightElement,
} from "@chakra-ui/react";

interface ISourceData {
  source: string;
}
interface ISourceProps {
  index: number;
  value: ISourceData;
  remove: any;
  register: any;
}

export const SourcesInput = ({
  register,
  remove,
  value,
  index,
}: ISourceProps) => {
  return (
    <Flex>
      <InputGroup size="md">
        <InputLeftAddon>{index}.</InputLeftAddon>

        <Input {...register()} />

        <InputRightElement>
          <IconButton
            aria-label={`remove source ${index}`}
            icon={<DeleteIcon />}
            onClick={() => remove(index)}
          />
        </InputRightElement>
      </InputGroup>
    </Flex>
  );
};
