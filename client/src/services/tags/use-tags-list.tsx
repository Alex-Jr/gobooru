import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { parseQueryString } from "shared/functions/url/parse-query-string";
import { APITagsList } from "shared/types/services/tags/APITagsList";

import { BASE_URL } from "../BASE_URL";

interface IUseTagsList {
  [index: string]: string;
  search: string;
  page: string;
  page_size: string;
}

export const useTagsLists = (
  queryData: IUseTagsList
): APITagsList & { totalPages: number } => {
  const { data } = useQuery({
    queryKey: ["tags", queryData],
    queryFn: async () => {
      const { data } = await axios<APITagsList>({
        method: "GET",
        url: `${BASE_URL}/tags?${parseQueryString(queryData)}`,
      });

      return data;
    },
  });

  if (!data)
    return {
      tags: [],
      count: 0,
      totalPages: 0,
    };

  return {
    tags: data.tags,
    count: data.count,
    totalPages: Math.ceil(data.count / +queryData.page_size),
  };
};
