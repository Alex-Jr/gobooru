import { SearchIcon } from "@chakra-ui/icons";
import { Flex, FormControl, FormLabel, IconButton } from "@chakra-ui/react";
import { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

import { useTagsLists } from "services/tags/use-tags-list";
import { InputWithSuggestions } from "shared/components/forms/input-with-suggestions";

interface IFormData {
  search: string;
}
export interface ISidebarProps {
  defaultValues: IFormData;
  onSubmit: (formData: IFormData) => void;
}

export const Sidebar = ({ defaultValues, onSubmit }: ISidebarProps) => {
  const { t } = useTranslation();

  const { register, control, handleSubmit, reset } = useForm({
    defaultValues,
  });

  const [search, setSearch] = useState(defaultValues.search);

  const { tags } = useTagsLists({
    page: "1",
    page_size: "10",
    search: search,
  });

  const tagSuggestions = tags.map((tag) => ({
    value: tag.id,
  }));

  return (
    <Flex
      as={"form"}
      bgColor={"gray.700"}
      gap={4}
      h={"fit-content"}
      justifyContent={"flex-start"}
      p={4}
      w={"fit-content"}
      onSubmit={handleSubmit(onSubmit)}
    >
      <FormControl w={"400px"}>
        <FormLabel htmlFor="search">{t("glossary.search")}</FormLabel>

        <Flex gap={2} mt={2}>
          <Controller
            name="search"
            control={control}
            rules={{ required: true }}
            render={({ field }) => (
              <InputWithSuggestions
                {...field}
                suggestions={tagSuggestions}
                onTypeEnd={(value) => {
                  const words = value.split(" ");

                  setSearch(words[words.length - 1]);
                }}
              />
            )}
          />

          <IconButton aria-label="search" type="submit" size={"sm"}>
            <SearchIcon />
          </IconButton>
        </Flex>
      </FormControl>
    </Flex>
  );

  // TODO: easy mode
};
