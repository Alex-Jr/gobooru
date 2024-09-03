import { ChevronLeftIcon, ChevronRightIcon } from "@chakra-ui/icons";
import { HStack, IconButton, Link, VStack } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

import { usePostsList } from "services/posts/use-posts-list";
import { useQueryString } from "shared/hooks/use-query-string";

export const QueryPagination = ({ postId }: { postId: number }) => {
  const {
    asString: { search = "", page = "1" },
  } = useQueryString({
    asString: ["search", "page"],
  });

  if (!search || search.length === 0) return <></>;

  // TODO: avoid preloading previous and next page if not needed

  // preloads previous page
  const { posts: previousPosts } = usePostsList({
    page_size: "20",
    search,
    page: (parseInt(page) - 1).toString(),
  });

  // this need to be equal to post-list-page
  const { posts } = usePostsList({
    page_size: "20",
    search,
    page,
  });

  // preloads next page
  const { posts: nextPosts } = usePostsList({
    page_size: "20",
    search,
    page: (parseInt(page) + 1).toString(),
  });

  const currentIndex = posts.findIndex((post) => post.id === postId);
  const previous = posts[currentIndex - 1] || previousPosts.at(-1);
  const next = posts[currentIndex + 1] || nextPosts[0];

  return (
    <VStack width={"100%"}>
      <HStack>
        <IconButton
          aria-label={"next"}
          as={RouterLink}
          colorScheme={previous && previous?.id > postId ? "orange" : "gray"}
          fontWeight={previous && previous?.id > postId ? "bold" : "normal"}
          pointerEvents={previous && previous?.id > postId ? "all" : "none"}
          to={{
            pathname: "/posts/" + previous?.id,
            search: `?search=${search}&page=${
              parseInt(page) -
              (previous?.id === previousPosts.at(-1)?.id ? 1 : 0)
            }`,
          }}
          icon={<ChevronLeftIcon />}
          size={"sm"}
        />

        <Link
          mx={5}
          fontSize={"xl"}
          as={RouterLink}
          to={{
            pathname: `/posts`,
            search: `?search=${search}&page=${page}`,
          }}
        >
          {search}
        </Link>

        <IconButton
          aria-label={"next"}
          as={RouterLink}
          colorScheme={next ? "orange" : "gray"}
          fontWeight={next ? "bold" : "normal"}
          pointerEvents={next ? "all" : "none"}
          to={{
            pathname: "/posts/" + next?.id,
            search: `?search=${search}&page=${
              parseInt(page) + (next?.id === nextPosts.at(0)?.id ? 1 : 0)
            }`,
          }}
          icon={<ChevronRightIcon />}
          size={"sm"}
        />
      </HStack>
    </VStack>
  );
};
