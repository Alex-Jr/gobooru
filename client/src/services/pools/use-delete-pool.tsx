import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

interface IUseDeletePool {
  id: number;
}

export const useDeletePool = () => {
  return useMutation<void, Error, IUseDeletePool>(["delete", "pool"], {
    mutationFn: (data) => {
      return axios({
        method: "DELETE",
        baseURL: `${BASE_URL}/pools/${data.id}`,
      });
    },
  });
};
