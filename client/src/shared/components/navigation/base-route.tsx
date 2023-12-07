import { Box, Flex } from "@chakra-ui/react";
import { ReactNode } from "react";

import { NavBar } from "./nav-bar";

interface IBaseRoute {
  children: ReactNode;
}

export const BaseRoute = ({ children }: IBaseRoute) => {
  return (
    <Flex direction={"column"} h={"100vh"}>
      <NavBar />
      <Box p={4} flex={1} overflowY={"auto"}>
        {children}
      </Box>
    </Flex>
  );
};
