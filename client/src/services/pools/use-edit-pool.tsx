import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

interface IUseEditPool {
  id: number;
  custom: string[];
  description: string;
  name: string;
  posts: number[];
}

export const useEditPool = <T,>() => {
  return useMutation<void, Error, IUseEditPool>(["edit", "pool"], {
    mutationFn: (data) => {
      return axios({
        method: "PATCH",
        baseURL: `${BASE_URL}/pools/${data.id}`,
        data: {
          custom: data.custom,
          description: data.description,
          name: data.name,
          posts: data.posts,
        },
      });
    },
  });
};
