import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { APIPool } from "shared/types/services/pools/APIPool";

import { BASE_URL } from "../BASE_URL";

interface IUsePool {
  id: string;
}

export const usePool = ({ id }: IUsePool): APIPool["pool"] | undefined => {
  const { data } = useQuery({
    queryKey: ["pools", id],
    queryFn: async () => {
      const { data } = await axios<APIPool>({
        method: "GET",
        url: `${BASE_URL}/pools/${id}`,
      });

      return data;
    },
  });

  if (!data) return undefined;

  return data.pool;
};
