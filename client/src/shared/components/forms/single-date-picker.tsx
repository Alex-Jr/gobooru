import { CalendarIcon } from "@chakra-ui/icons";
import {
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
} from "@chakra-ui/react";
import { format } from "date-fns";
import { forwardRef, useRef } from "react";

interface SingleDatePickerProps {
  value: Date;
  name: string;
  onChange: (...event: any[]) => void;
  disabled?: boolean;
  onBlur: () => void;
}

export const SingleDatePicker = forwardRef<
  HTMLInputElement,
  SingleDatePickerProps
>(function SingleDatePicker(props, ref) {
  const innerRef = useRef<HTMLInputElement | null>(null);

  return (
    <InputGroup size="sm">
      <Input
        onReset={() => {
          props.onChange(undefined);
        }}
        disabled={props.disabled}
        value={format(props.value, "yyyy-MM-dd")}
        onBlur={props.onBlur}
        onChange={(e) => {
          props.onChange(new Date(e.target.value));
        }}
        type={"date"}
        ref={(element) => {
          innerRef.current = element;

          if (typeof ref === "function") {
            ref(element);
          } else {
            innerRef.current = element;
          }
        }}
      />

      <InputRightElement>
        <IconButton
          size="sm"
          aria-label={`show ${props.name} picker`}
          icon={<CalendarIcon />}
          onReset={() => {
            if (innerRef.current) {
              innerRef.current.value = "";
            }
          }}
          onClick={() => {
            if (innerRef.current) {
              innerRef.current.showPicker();
            }
          }}
        />
      </InputRightElement>
    </InputGroup>
  );
});
