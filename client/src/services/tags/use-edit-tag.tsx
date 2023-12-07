import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

export const useEditTag = <T,>() => {
  const mutationKey = "delete-tag";

  return useMutation<
    unknown,
    unknown,
    {
      id: string;
      categoryId: string;
      description: string;
    },
    unknown
  >([mutationKey], {
    mutationFn: ({ id, categoryId, description }) => {
      console.log(description);
      return axios({
        method: "PATCH",
        baseURL: `${BASE_URL}/tags/${id}`,
        data: {
          categoryId,
          description,
        },
      });
    },
  });
};
