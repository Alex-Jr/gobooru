import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { APIPost } from "shared/types/services/posts/APIPost";

import { BASE_URL } from "../BASE_URL";

interface IUsePost {
  id: string;
}

export const usePost = ({ id }: IUsePost): APIPost["post"] | undefined => {
  const { data } = useQuery({
    queryKey: ["posts", id],
    queryFn: async () => {
      if (!id) return { post: undefined };

      const { data } = await axios<APIPost>({
        method: "GET",
        url: `${BASE_URL}/posts/${id}`,
      });

      return data;
    },
  });

  if (!data) return undefined;

  return data.post;
};
