import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { APITag } from "shared/types/services/tags/APITag";

import { BASE_URL } from "../BASE_URL";

interface IUsePostParams {
  id: string;
}

export const useTag = ({ id }: IUsePostParams): APITag["tag"] | undefined => {
  const { data } = useQuery({
    queryKey: ["tags", id],
    queryFn: async () => {
      const { data } = await axios<APITag>({
        method: "GET",
        url: `${BASE_URL}/tags/${id}`,
      });

      return data;
    },
  });

  if (!data) return undefined;

  return data.tag;
};
