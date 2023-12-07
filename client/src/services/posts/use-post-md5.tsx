import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { APIPost } from "shared/types/services/posts/APIPost";

import { BASE_URL } from "../BASE_URL";

interface IUsePost {
  md5: string;
}

export const usePostMD5 = ({ md5 }: IUsePost): APIPost["post"] | undefined => {
  const { data } = useQuery({
    queryKey: ["posts", md5],
    queryFn: async () => {
      if (!md5) return { post: undefined };

      const { data } = await axios<APIPost>({
        method: "GET",
        url: `${BASE_URL}/posts/hash/${md5}`,
        validateStatus(s) {
          // TODO: remove 500 after server starts returning 404
          return [200, 404, 500].includes(s);
        },
      });

      return data;
    },
  });

  if (!data || !data.post) return undefined;

  return data.post;
};
