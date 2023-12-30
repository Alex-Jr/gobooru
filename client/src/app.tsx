import "./app.css";
import "./i18n";

import { Box, ChakraProvider, Spinner, useBoolean } from "@chakra-ui/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { useEffect, useState } from "react";

import { UserContext } from "shared/context/userContext";
import { APIUser } from "shared/types/services/user/APIUser";

import { MainRouter } from "./routes";
import theme from "./theme";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 60 * 1000,
    },
  },
});

function Loading() {
  return (
    <Box w={"100vw"} h={"100vh"}>
      <Spinner
        position={"absolute"}
        top={"50%"}
        left={"50%"}
        transform={"translate(-50%, -50%)"}
        size={"lg"}
      />
    </Box>
  );
}

export const App = () => {
  const [user, setUser] = useState<APIUser | undefined>();

  const [loading, { off: setLoadingFalse }] = useBoolean(true);

  useEffect(() => {
    // setTimeout(() => {
    setLoadingFalse();
    // }, 1000);
  });

  // useEffect(() => {
  //   authCheck().then(setUser);
  // }, []);

  return (
    <ChakraProvider theme={theme}>
      {/* this prevent first load flicking */}
      {loading ? (
        <Loading />
      ) : (
        <QueryClientProvider client={queryClient}>
          <UserContext.Provider value={{ user, setUser }}>
            <MainRouter />
          </UserContext.Provider>
          <ReactQueryDevtools initialIsOpen={false} />
        </QueryClientProvider>
      )}
    </ChakraProvider>
  );
};
