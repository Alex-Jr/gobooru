import { Box, Text, VStack } from "@chakra-ui/react";
import { useTranslation } from "react-i18next";

import { PostCount } from "./components/post-count";
import { SearchForm } from "./components/search-form";

export const HomePage = () => {
  const { t } = useTranslation();

  return (
    <Box
      h={"100%"}
      w={"500px"}
      m={"auto"}
      bgImage={
        "https://i0.wp.com/imagensemoldes.com.br/wp-content/uploads/2021/08/Imagem-Raccoon-PNG.png"
      }
      bgPos={"center"}
      bgSize={"contain"}
      bgRepeat={"no-repeat"}
      pt={"50vh"}
    >
      <VStack
        zIndex={1}
        gap={4}
        py={2}
        w={"300px"}
        m={"auto"}
        rounded="md"
        boxShadow={"lg"}
        backdropFilter={"blur(15px)"}
      >
        <Text fontSize={"6xl"}>{t("homepage.title")}</Text>

        <SearchForm />

        <PostCount />
      </VStack>
    </Box>
  );
};
