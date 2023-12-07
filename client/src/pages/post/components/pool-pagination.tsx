import { ChevronLeftIcon, ChevronRightIcon } from "@chakra-ui/icons";
import { HStack, IconButton, Link, VStack } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

import { APIPost } from "shared/types/services/posts/APIPost";

const PoolPaginationItem = ({
  postId,
  pool,
}: {
  postId: number;
  pool: {
    id: number;
    name: string;
    post_ids: number[];
    custom: string[];
    description: string;
  };
}) => {
  const index = pool.post_ids.findIndex((post_id) => post_id === postId);

  const next = pool.post_ids[index + 1];
  const previous = pool.post_ids[index - 1];

  return (
    <HStack key={`pool-${pool.id}`}>
      <IconButton
        aria-label={"next"}
        as={RouterLink}
        colorScheme={previous ? "orange" : "gray"}
        fontWeight={previous ? "bold" : "normal"}
        pointerEvents={previous ? "all" : "none"}
        to={"/posts/" + previous}
        key={`pagination-${previous}`}
        icon={<ChevronLeftIcon />}
        size={"sm"}
      />

      <Link mx={5} fontSize={"xl"} as={RouterLink} to={`/pools/${pool.id}`}>
        {pool.name}
      </Link>

      <IconButton
        aria-label={"next"}
        as={RouterLink}
        colorScheme={next ? "orange" : "gray"}
        fontWeight={next ? "bold" : "normal"}
        pointerEvents={next ? "all" : "none"}
        to={"/posts/" + next}
        key={`pagination-${next}`}
        icon={<ChevronRightIcon />}
        size={"sm"}
      />
    </HStack>
  );
};

interface IPoolPaginationProps {
  post: APIPost["post"];
}

export const PoolPagination = (props: IPoolPaginationProps) => {
  if (!props.post.pools) return <></>;

  return (
    <VStack width={"100%"}>
      {props.post.pools.map((pool) => (
        <PoolPaginationItem
          key={`pool-pagination-${pool.id}`}
          pool={pool}
          postId={props.post.id}
        />
      ))}
    </VStack>
  );
};
