import {
  Input,
  List,
  ListItem,
  Popover,
  PopoverBody,
  PopoverContent,
  PopoverTrigger,
  useDisclosure,
} from "@chakra-ui/react";
import React, { ChangeEvent, forwardRef, useRef, useState } from "react";
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
  const { isOpen, onClose, onOpen } = useDisclosure();

  const inputRef = useRef<HTMLInputElement | null>(null);

  const [highlightedIndex, setHighlightedIndex] = useState(-1);

  const onTypeEndDebounced = useDebounce(props.onTypeEnd, 500);

  const suggestions = props.suggestions.filter((suggestion) => {
    return !props.value.includes(suggestion.value);
  });

  const handleClose = () => {
    props.onBlur();
    onClose();
    setHighlightedIndex(-1);
  };

  const handleOnChange = (e: ChangeEvent<HTMLInputElement>) => {
    console.log("a");
    props.onChange(e);
    onTypeEndDebounced(e.target.value);
  };

  const handleSuggestionPick = () => {
    const words = props.value.split(" ");

    props.onChange(
      words
        .filter((w: string) => w !== "")
        .slice(
          0,
          words[words.length - 1] === "" ? words.length : words.length - 1
        )
        .concat(suggestions[highlightedIndex].value)
        .join(" ")
    );

    setHighlightedIndex(-1);

    props.onTypeEnd("");
  };

  return (
    <>
      <Popover
        initialFocusRef={inputRef}
        isOpen={isOpen}
        returnFocusOnClose={false}
      >
        <>
          <PopoverTrigger>
            <Input
              size={"sm"}
              name={props.name}
              autoComplete="off"
              value={props.value}
              ref={(el) => {
                inputRef.current = el;
                if (ref) {
                  if (typeof ref === "function") {
                    ref(inputRef.current);
                  }
                }
              }}
              onChange={handleOnChange}
              onKeyDown={(e) => {
                if (e.key === "Escape") {
                  handleClose();
                  return e.preventDefault();
                }
                if (e.key === "ArrowDown") {
                  setHighlightedIndex((prevIndex) =>
                    Math.min(prevIndex + 1, suggestions.length - 1)
                  );
                  return e.preventDefault();
                }
                if (e.key === "ArrowUp") {
                  setHighlightedIndex((prevIndex) =>
                    Math.max(prevIndex - 1, -1)
                  );
                  return e.preventDefault();
                }
                if (e.key === "Enter") {
                  if (highlightedIndex === -1) {
                    return handleClose();
                  }

                  handleSuggestionPick();

                  return e.preventDefault();
                }
              }}
              onBlur={() => {
                handleClose();
              }}
              onFocus={() => {
                onOpen();
              }}
            />
          </PopoverTrigger>
          {suggestions.length > 0 && (
            <PopoverContent
              width={inputRef.current?.clientWidth || "100%"}
              onFocus={() => {
                inputRef.current?.focus();
              }}
            >
              <PopoverBody p={0}>
                <List>
                  {suggestions.map((suggestion, index) => (
                    <ListItem
                      fontSize={"sm"}
                      key={index}
                      bg={
                        highlightedIndex === index ? "gray.800" : "transparent"
                      }
                      p={2}
                      onMouseEnter={(e) => {
                        setHighlightedIndex(index);
                      }}
                      onMouseDown={(e) => {
                        handleSuggestionPick();
                      }}
                    >
                      {suggestion.value}
                    </ListItem>
                  ))}
                </List>
              </PopoverBody>
            </PopoverContent>
          )}
        </>
      </Popover>
    </>
  );
});
