import {
  HStack,
  Input,
  InputProps,
  Tag,
  TagCloseButton,
  TagLabel,
  Tooltip,
} from "@chakra-ui/react";
import { KeyboardEvent, forwardRef } from "react";

interface InputWithTagProps extends InputProps {
  value: string[];
  name: string;
  onChange: (...event: any[]) => void;
  disabled?: boolean;
  onBlur: () => void;
}

export const InputWithTag = forwardRef<HTMLInputElement, InputWithTagProps>(
  function InputWithTag({ value, name, onChange, onBlur, ...rest }, ref) {
    const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
      // enter or space will save current input value in the array
      if (e.key === "Enter" || e.key === " ") {
        if (e.currentTarget.value.length > 0) {
          if (value.includes(e.currentTarget.value)) {
            e.preventDefault();

            return;
          }

          onChange([...value, e.currentTarget.value.replaceAll(" ", "")]);

          e.currentTarget.value = "";

          e.preventDefault();
        }
      }
    };

    return (
      <>
        <Input
          {...rest}
          ref={ref}
          name={"inner" + name}
          onBlur={onBlur}
          autoComplete="off"
          onKeyDown={handleKeyDown}
        />

        {value.length > 0 && (
          <HStack wrap={"wrap"} mt={2}>
            {value.map((tag, index) => (
              <Tooltip
                key={index.toString() + tag}
                label={tag.length > 17 ? tag : undefined}
              >
                <Tag w={"fit-content"}>
                  <TagLabel>{tag}</TagLabel>
                  <TagCloseButton
                    onClick={() => {
                      const newArr = [...value];

                      newArr.splice(index, 1);

                      onChange(newArr);
                    }}
                  />
                </Tag>
              </Tooltip>
            ))}
          </HStack>
        )}
      </>
    );
  }
);
