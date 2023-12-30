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

interface SidebarProps {
  defaultValues: IFormData;
  onSubmit: (formData: IFormData) => void;
}

export const Sidebar = ({ defaultValues, onSubmit }: SidebarProps) => {
  const { t } = useTranslation();

  const { register, control, handleSubmit } = useForm({ defaultValues });

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
  // return (
  //   <Flex
  //     as={"form"}
  //     justifyContent={"flex-start"}
  //     direction={"column"}
  //     w={275}
  //     bgColor={useColorModeValue("lightgray", "gray.700")}
  //     p={4}
  //     h={"fit-content"}
  //     maxH={"85vh"}
  //     display={{ sm: "none", lg: "flex" }}
  //     gap={2}
  //   >
  //     <FormControl>
  //       <FormLabel m={0} htmlFor="tags">
  //         {t("glossary.name")}
  //       </FormLabel>

  //       <Input {...register("id")} autoComplete="off" size={"sm"} />
  //     </FormControl>

  //     <FormControl>
  //       <FormLabel htmlFor="fromDate">{t("glossary.fromDate")}</FormLabel>

  //       <Controller
  //         control={control}
  //         name={"fromDate"}
  //         render={({ field }) => <SingleDatePicker {...field} />}
  //       />
  //     </FormControl>

  //     <FormControl>
  //       <FormLabel htmlFor="toDate">{t("glossary.toDate")}</FormLabel>

  //       <Controller
  //         control={control}
  //         name={"toDate"}
  //         render={({ field }) => <SingleDatePicker {...field} />}
  //       />
  //     </FormControl>

  //     <FormControl>
  //       <FormLabel htmlFor="orderBy">{t("glossary.orderBy")}</FormLabel>

  //       <Select {...register("orderBy")} size={"sm"}>
  //         <option value={"id"}>{t("glossary.id")}</option>
  //         <option value={"usageCount"}>{t("glossary.count")}</option>
  //       </Select>
  //     </FormControl>

  //     <HStack mt={2}>
  //       <Button
  //         alignSelf={"flex-end"}
  //         w={"100%"}
  //         colorScheme="yellow"
  //         rightIcon={<RepeatIcon />}
  //         size={"sm"}
  //         type={"reset"}
  //       >
  //         {t("glossary.clear")}
  //       </Button>

  //       <Button
  //         alignSelf={"flex-end"}
  //         w={"100%"}
  //         colorScheme="green"
  //         rightIcon={<SearchIcon />}
  //         size={"sm"}
  //         type="submit"
  //       >
  //         {t("glossary.search")}
  //       </Button>
  //     </HStack>
  //   </Flex>
  // );
};
