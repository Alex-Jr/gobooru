import { DeleteIcon } from "@chakra-ui/icons";
import {
  Flex,
  IconButton,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightElement,
} from "@chakra-ui/react";

interface ICustomData {
  custom: string;
}

interface ICustomProps {
  index: number;
  value: ICustomData;
  remove: any;
  register: any;
}

export const CustomInput = ({
  register,
  remove,
  value,
  index,
}: ICustomProps) => {
  return (
    <Flex>
      <InputGroup size="md">
        <InputLeftAddon>{index}°</InputLeftAddon>

        <Input {...register()} />

        <InputRightElement>
          <IconButton
            aria-label={`remove custom ${index}`}
            icon={<DeleteIcon />}
            onClick={() => remove(index)}
          />
        </InputRightElement>
      </InputGroup>
    </Flex>
  );
};
