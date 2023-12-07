import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

interface IUseCreatePool {
  custom: string[];
  name: string;
  description: string;
  posts: number[];
}

interface IUseCreatePoolResponse {
  pool: {
    id: number;
  };
}

export const useCreatePool = () => {
  return useMutation<IUseCreatePoolResponse, Error, IUseCreatePool>(
    ["edit", "post"],
    {
      mutationFn: async (data) => {
        const { data: axiosData } = await axios({
          method: "POST",
          url: `${BASE_URL}/pools`,
          data,
          withCredentials: true,
        });

        return axiosData;
      },
    }
  );
};
