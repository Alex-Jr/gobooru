import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "services/BASE_URL";

export const useLogout = () => {
  return useMutation<void, Error, void>(["logout"], {
    mutationFn: async () => {
      // TODO: users are not implemented yet
      if (1 > 0) return;

      return axios({
        method: "POST",
        url: `${BASE_URL}/auth/logout`,
        withCredentials: true,
      });
    },
  });
};
