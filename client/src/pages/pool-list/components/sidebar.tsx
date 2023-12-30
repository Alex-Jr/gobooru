import { SearchIcon } from "@chakra-ui/icons";
import {
  Flex,
  FormControl,
  FormLabel,
  IconButton,
  Input,
} from "@chakra-ui/react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";

export interface IFormInput {
  search: string;
}

export const Sidebar = (props: {
  initialValues: IFormInput;
  onSubmit: (formData: IFormInput) => void;
}) => {
  const { t } = useTranslation();

  const { register, handleSubmit, reset, control } = useForm<IFormInput>({
    defaultValues: props.initialValues,
  });

  const onSubmit: SubmitHandler<IFormInput> = (formData) => {
    props.onSubmit(formData);
  };

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
  //     bgColor={"gray.700"}
  //     direction={"column"}
  //     gap={2}
  //     h={"fit-content"}
  //     justifyContent={"flex-start"}
  //     maxH={"85vh"}
  //     mx={"auto"}
  //     p={4}
  //     w={{ sm: "100%", lg: 275 }}
  //     onSubmit={handleSubmit(onSubmit)}
  //   >
  //     <FormControl>
  //       <FormLabel htmlFor="name">{t("glossary.name")}</FormLabel>

  //       <Input {...register("name")} size="sm" />
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
  //       <FormLabel htmlFor="orderBy">Order by</FormLabel>
  //       <Select {...register("orderBy")} size={"sm"}>
  //         <option>id</option>
  //         <option>name</option>
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
  //         onClick={() =>
  //           reset({
  //             fromDate: undefined,
  //           })
  //         }
  //       >
  //         Clear
  //       </Button>

  //       <Button
  //         alignSelf={"flex-end"}
  //         w={"100%"}
  //         colorScheme="green"
  //         rightIcon={<SearchIcon />}
  //         size={"sm"}
  //         type="submit"
  //       >
  //         Search
  //       </Button>
  //     </HStack>
  //   </Flex>
  // );
};
