import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { parseQueryString } from "shared/functions/url/parse-query-string";
import { APIPostList } from "shared/types/services/posts/APIPostList";

import { BASE_URL } from "../BASE_URL";

interface IUsePostList {
  [index: string]: string | string[];
  search: string;
  page: string;
  page_size: string;
}

export const usePostsList = (
  queryData: IUsePostList
): APIPostList & { totalPages: number } => {
  const { data } = useQuery({
    queryKey: ["posts", queryData],
    queryFn: async () => {
      if (!queryData.search.includes("status")) {
        queryData.search += " status:ACTIVE";
      }

      const { data } = await axios<APIPostList>({
        method: "GET",
        url: `${BASE_URL}/posts?${parseQueryString(queryData)}`,
      });

      return data;
    },
  });

  if (!data) {
    return {
      posts: [],
      count: 0,
      totalPages: 0,
    };
  }

  return {
    posts: data.posts,
    count: data.count,
    totalPages: Math.ceil(data.count / +queryData.page_size),
  };
};
