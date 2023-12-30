import {
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
} from "@chakra-ui/react";
import { forwardRef, useState } from "react";
import { ChangeHandler } from "react-hook-form";
import { FaEye, FaEyeSlash } from "react-icons/fa";

interface PasswordInputProps {
  name: string;
  placeholder: string;
  onChange: ChangeHandler;
  onBlur: ChangeHandler;
}

export const PasswordInput = forwardRef<HTMLInputElement, PasswordInputProps>(
  function PasswordInput(props, ref) {
    const [isHidden, setIsHidden] = useState(true);

    return (
      <InputGroup size="md">
        <Input type={isHidden ? "password" : "text"} {...props} ref={ref} />

        <InputRightElement>
          <IconButton
            aria-label={isHidden ? `show ${props.name}` : `hide ${props.name}`}
            icon={isHidden ? <FaEyeSlash /> : <FaEye />}
            onClick={() => setIsHidden(!isHidden)}
          />
        </InputRightElement>
      </InputGroup>
    );
  }
);
