import { Flex, Link, Td, Tr } from "@chakra-ui/react";
import { useTranslation } from "react-i18next";
import { Link as RouterLink } from "react-router-dom";

import { useTagCategoryList } from "services/tag-categories/use-tag-category";
import { useTagsLists } from "services/tags/use-tags-list";
import { GenericTable } from "shared/components/lists/generic-table";
import { Pagination } from "shared/components/lists/pagination";
import { useQueryString } from "shared/hooks/use-query-string";

import { Sidebar } from "./components/sidebar";

export const TagListPage = () => {
  const {
    asString: { search = "", page = "1" },
    setSearchParams,
  } = useQueryString({
    asString: ["search", "page"],
  });

  const { tagCategoriesObj } = useTagCategoryList();

  const { tags, count, totalPages } = useTagsLists({
    search,
    page,
    pageSize: "23",
  });

  const { t } = useTranslation();

  return (
    <Flex direction={"column"} h={"100%"}>
      <Flex gap={4}>
        <Sidebar
          defaultValues={{
            search,
          }}
          onSubmit={(formData) => {
            setSearchParams({
              ...formData,
            });
          }}
        />

        <GenericTable
          size={"sm"}
          headers={[
            t("glossary.tag"),
            t("glossary.usage"),
            t("glossary.category"),
          ]}
          rows={tags}
          renderRow={(row, index) => (
            <Tr key={`row-${index}`}>
              <Td>
                <Link
                  w={"100%"}
                  display={"block"}
                  as={RouterLink}
                  to={`/tags/${row.id}`}
                >
                  {row.id}
                </Link>
              </Td>
              <Td>{row.post_count}</Td>
              <Td color={tagCategoriesObj[row.category_id]}>
                {row.category_id}
              </Td>
            </Tr>
          )}
        />
      </Flex>

      <Pagination totalPages={totalPages} currentPage={parseInt(page)} />
    </Flex>
  );
};
