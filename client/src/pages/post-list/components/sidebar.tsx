import { SearchIcon } from "@chakra-ui/icons";
import {
  Flex,
  FormControl,
  FormLabel,
  IconButton,
  Input,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

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
          <Input {...register("search")} size={"sm"} />

          <IconButton aria-label="search" type="submit" size={"sm"}>
            <SearchIcon />
          </IconButton>
        </Flex>
      </FormControl>
    </Flex>
  );

  // TODO: easy mode
};
