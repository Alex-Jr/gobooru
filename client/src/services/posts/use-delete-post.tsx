import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

interface IUseDeletePost {
  id: number;
  type: "SOFT" | "HARD";
}

export const useDeletePost = () => {
  return useMutation<void, Error, IUseDeletePost>(["delete", "post"], {
    mutationFn: (data) => {
      if (data.type === "SOFT") {
        return axios({
          method: "DELETE",
          baseURL: `${BASE_URL}/post/${data.id}/file`,
        });
      }

      return axios({
        method: "DELETE",
        baseURL: `${BASE_URL}/posts/${data.id}`,
      });
    },
  });
};
