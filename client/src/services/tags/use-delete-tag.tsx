import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

export const useDeleteTag = <T,>() => {
  const mutationKey = "delete-tag";

  return useMutation<
    unknown,
    unknown,
    {
      id: string;
    },
    unknown
  >([mutationKey], {
    mutationFn: (data) => {
      return axios({
        method: "DELETE",
        baseURL: `${BASE_URL}/tags/${data.id}`,
      });
    },
  });
};
