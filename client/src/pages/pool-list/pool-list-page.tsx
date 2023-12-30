import { Flex } from "@chakra-ui/react";

import { usePoolList } from "services/pools/use-pool-list";
import { GenericGrid } from "shared/components/lists/generic-list";
import { Pagination } from "shared/components/lists/pagination";
import { PoolItem } from "shared/components/lists/pool-item";
import { useQueryString } from "shared/hooks/use-query-string";

import { Sidebar } from "./components/sidebar";

export const PoolListPage = () => {
  const {
    asString: { search = "", page = "1" },
    setSearchParams,
  } = useQueryString({
    asString: ["search", "page"],
  });

  const { pools, count, totalPages } = usePoolList({
    search,
    page,
    pageSize: "12",
  });

  return (
    <Flex direction={"column"} minH={"100%"}>
      <Flex gap={4} direction={{ sm: "column", lg: "column" }}>
        <Sidebar
          initialValues={{
            search: search || "",
          }}
          onSubmit={(formData) => {
            setSearchParams({ ...formData });
          }}
        />

        <GenericGrid
          w="250px"
          items={pools}
          renderItem={(pool, index) => (
            <PoolItem key={`pool-item-${index}`} pool={pool} />
          )}
        />
      </Flex>

      <Pagination totalPages={totalPages} currentPage={parseInt(page, 10)} />
    </Flex>
  );
};
