import { Flex } from "@chakra-ui/react";
import { useState } from "react";

import { usePostsList } from "services/posts/use-posts-list";
import { GenericGrid } from "shared/components/lists/generic-list";
import { Pagination } from "shared/components/lists/pagination";
import { PostItem } from "shared/components/lists/post-item";
import { useQueryString } from "shared/hooks/use-query-string";

import { Sidebar } from "./components/sidebar";

export const PostListPage = () => {
  const {
    asString: { search = "", page = "1" },
    setSearchParams,
  } = useQueryString({
    asString: ["search", "page"],
  });

  const [pageSize, setPageSize] = useState("20");

  const { posts, count, totalPages } = usePostsList({
    page_size: pageSize,
    search,
    page,
  });

  return (
    <Flex direction={"column"} minH={"100%"}>
      <Flex gap={4} direction={"column"}>
        <Sidebar
          defaultValues={{
            search,
          }}
          onSubmit={({ search }) => {
            setSearchParams({
              search,
              page,
            });
          }}
        />

        <GenericGrid
          w="250px"
          items={posts}
          renderItem={(post, index) => (
            <PostItem key={`pool-item-${index}`} post={post} />
          )}
        />
      </Flex>

      <Pagination totalPages={totalPages} currentPage={parseInt(page)} />
    </Flex>
  );
};
