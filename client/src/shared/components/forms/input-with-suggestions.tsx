import {
  Input,
  List,
  ListItem,
  Popover,
  PopoverBody,
  PopoverContent,
  PopoverTrigger,
} from "@chakra-ui/react";
import React, { forwardRef, useRef, useState } from "react";
import { ControllerRenderProps } from "react-hook-form";

import { useDebounce } from "shared/hooks/use-debounce";

interface InputWithSuggestionsProps extends ControllerRenderProps {
  suggestions: { value: string; text?: string }[];
  onTypeEnd: (value: string) => void;
}

export const InputWithSuggestions = forwardRef<
  HTMLInputElement,
  InputWithSuggestionsProps
>(function InputWithSuggestions(props, ref) {
  const inputRef = useRef<HTMLInputElement | null>(null);

  const [highlightedIndex, setHighlightedIndex] = useState(-1);

  const onTypeEndDebounced = useDebounce(props.onTypeEnd, 500);

  const suggestions = props.suggestions.filter((suggestion) => {
    return !props.value.includes(suggestion.value);
  });

  return (
    <>
      <Popover initialFocusRef={inputRef}>
        {({ isOpen, onClose }) => (
          <>
            <PopoverTrigger>
              <Input
                name={props.name}
                autoComplete="off"
                size={"sm"}
                ref={(el) => {
                  inputRef.current = el;
                  if (ref) {
                    if (typeof ref === "function") {
                      ref(inputRef.current);
                    }
                  }
                }}
                value={props.value}
                onChange={(e) => {
                  props.onChange(e);

                  onTypeEndDebounced(e.target.value);
                }}
                onKeyDown={(e) => {
                  // console.log("key down", e.key);
                  if (e.key === "ArrowDown") {
                    setHighlightedIndex((prevIndex) =>
                      Math.min(prevIndex + 1, suggestions.length - 1)
                    );
                    e.preventDefault();
                  } else if (e.key === "ArrowUp") {
                    setHighlightedIndex((prevIndex) =>
                      Math.max(prevIndex - 1, -1)
                    );
                    e.preventDefault();
                  } else if (e.key === "Enter") {
                    if (highlightedIndex === -1) {
                      onClose();
                      return;
                    }

                    const words = props.value.split(" ");

                    console.log(
                      words
                        .slice(0, words.length - 1)
                        .concat(suggestions[highlightedIndex].value)
                        .join(" ")
                        .trim()
                    );

                    props.onChange(
                      words
                        .filter((w: string) => w !== "")
                        .slice(
                          0,
                          words[words.length - 1] === ""
                            ? words.length
                            : words.length - 1
                        )
                        .concat(suggestions[highlightedIndex].value)
                        .join(" ")
                    );

                    setHighlightedIndex(-1);

                    e.preventDefault();
                  }
                }}
                onBlur={(e) => {
                  onClose();
                }}
              />
            </PopoverTrigger>
            {suggestions.length > 0 && (
              <PopoverContent justifySelf={"start"}>
                <PopoverBody p={0}>
                  <List>
                    {suggestions.map((suggestion, index) => (
                      <ListItem
                        fontSize={"sm"}
                        key={index}
                        bg={
                          highlightedIndex === index
                            ? "gray.800"
                            : "transparent"
                        }
                        p={2}
                      >
                        {suggestion.value}
                      </ListItem>
                    ))}
                  </List>
                </PopoverBody>
              </PopoverContent>
            )}
          </>
        )}
      </Popover>
    </>
  );
});
