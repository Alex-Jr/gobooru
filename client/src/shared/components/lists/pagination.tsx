import { ChevronLeftIcon, ChevronRightIcon } from "@chakra-ui/icons";
import { Button, HStack, IconButton } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

import { useQueryString } from "shared/hooks/use-query-string";

interface IPagination {
  totalPages: number;
  currentPage: number;
  maxVisiblePages?: number;
}

export const Pagination = ({
  totalPages,
  currentPage,
  maxVisiblePages = 9,
}: IPagination) => {
  const { searchParams } = useQueryString({});

  searchParams.delete("page");

  const curSearch = searchParams.toString();

  if (totalPages <= 0) return <></>;

  const halfVisiblePage = Math.floor(maxVisiblePages / 2);
  let pagesToShow: number[] = [];

  if (totalPages <= maxVisiblePages) {
    // Total pages is smaller than maxVisiblePages -> show all pages
    pagesToShow = Array(totalPages)
      .fill(undefined)
      .map((_, i) => i + 1);
  } else if (currentPage <= halfVisiblePage) {
    // Current page is near the start -> show first maxVisiblePages
    pagesToShow = Array(maxVisiblePages)
      .fill(undefined)
      .map((_, i) => i + 1);
  } else if (currentPage > totalPages - halfVisiblePage) {
    // Current page is near the end -> show last maxVisiblePages
    pagesToShow = Array(maxVisiblePages)
      .fill(undefined)
      .map((_, i) => totalPages - maxVisiblePages + i + 1);
  } else {
    // Default case -> show middle maxVisiblePages
    const startPage = currentPage - halfVisiblePage;
    pagesToShow = Array(maxVisiblePages)
      .fill(undefined)
      .map((_, i) => startPage + i);
  }

  return (
    <HStack mt={"auto"} mx={"auto"}>
      <IconButton
        as={RouterLink}
        aria-label="Previous page"
        icon={<ChevronLeftIcon />}
        pointerEvents={currentPage === 1 ? "none" : "all"}
        to={{
          search: curSearch + `&page=${currentPage - 1}`,
        }}
      />

      {pagesToShow.map((page, index) => (
        <Button
          as={RouterLink}
          colorScheme={page === currentPage ? "orange" : "gray"}
          fontWeight={page === currentPage ? "bold" : "normal"}
          pointerEvents={page === currentPage ? "none" : "all"}
          to={{
            search: curSearch + `&page=${page}`,
          }}
          key={`pagination-${index}`}
        >
          {page}
        </Button>
      ))}

      <IconButton
        as={RouterLink}
        aria-label="Next page"
        icon={<ChevronRightIcon />}
        pointerEvents={currentPage === pagesToShow.at(-1) ? "none" : "all"}
        to={{
          search: curSearch + `&page=${currentPage + 1}`,
        }}
      />
    </HStack>
  );
};
