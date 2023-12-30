import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

interface IUseEditPost {
  custom: string[];
  description: string;
  id: number;
  rating: string;
  sources: string[];
  status: string;
  tags: string[];
}

export const useEditPost = <T,>() => {
  return useMutation<void, Error, IUseEditPost>(["edit", "post"], {
    mutationFn: (data) => {
      return axios({
        method: "PATCH",
        baseURL: `${BASE_URL}/posts/${data.id}`,
        data: {
          custom: data.custom,
          description: data.description,
          rating: data.rating,
          sources: data.sources,
          status: data.status,
          tags: data.tags,
        },
      });
    },
  });
};
