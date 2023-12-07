import { Input } from "@chakra-ui/react";
import React, { KeyboardEvent, forwardRef, useRef } from "react";

interface InputWithSuggestionsProps {
  // value: string;
  name: string;
  // onChange: (...event: any[]) => void;
  disabled?: boolean;
  onBlur: any;
  suggestions: { value: string; text: string }[];
  onTypeEnd: (value: string) => void;
  onKeyDown?: (e: KeyboardEvent<HTMLInputElement>) => void;
}

export const InputWithSuggestions = forwardRef<
  HTMLInputElement,
  InputWithSuggestionsProps
>(function InputWithSuggestions(props, ref) {
  const innerRef = useRef<HTMLInputElement | null>();

  let timeout: NodeJS.Timeout | null = null;

  return (
    <>
      <Input
        ref={(element) => {
          innerRef.current = element;

          if (typeof ref === "function") {
            ref(element);
          } else {
            innerRef.current = element;
          }
        }}
        onBlur={props.onBlur}
        name={props.name}
        autoComplete="off"
        position={"relative"}
        list={`${props.name}suggestions`}
        onKeyDown={(e) => {
          // debounce
          if (timeout) {
            clearTimeout(timeout);
          }

          timeout = setTimeout(() => {
            if (innerRef.current?.value)
              props.onTypeEnd(innerRef.current?.value);
          }, 1000);

          // prevent spamming when Enter or Space triggers an rerender
          if (e.key === "Enter" || e.key === " ") {
            clearTimeout(timeout);
          }

          if (props.onKeyDown) props.onKeyDown(e);
        }}
      />

      <datalist id={`${props.name}suggestions`}>
        {props.suggestions.map((s) => (
          <option
            key={`${props.name}-suggestion-${s.text}-${s.value}`}
            value={s.value}
          >
            {s.text}
          </option>
        ))}
      </datalist>
    </>
  );
});
