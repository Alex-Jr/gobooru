import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { parseQueryString } from "shared/functions/url/parse-query-string";
import { APIPoolList } from "shared/types/services/pools/APIPoolList";

import { BASE_URL } from "../BASE_URL";

interface IUsePostList {
  [index: string]: string;
  search: string;
  page: string;
  pageSize: string;
}

export const usePoolList = (
  queryData: IUsePostList
): APIPoolList & { totalPages: number } => {
  const { data } = useQuery({
    queryKey: ["pools", queryData],
    queryFn: async () => {
      const { data } = await axios<APIPoolList>({
        method: "GET",
        url: `${BASE_URL}/pools?${parseQueryString(queryData)}`,
      });

      return data;
    },
  });

  if (!data)
    return {
      pools: [],
      count: 0,
      totalPages: 0,
    };

  return {
    pools: data.pools,
    count: data.count,
    totalPages: Math.ceil(data.count / parseInt(queryData.pageSize)),
  };
};
