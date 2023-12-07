import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "services/BASE_URL";
import { APIUser } from "shared/types/services/user/APIUser";

interface IUseLoginMutation {
  name: string;
  password: string;
}

export const useLoginMutation = () => {
  return useMutation<APIUser, Error, IUseLoginMutation>(["login"], {
    mutationFn: async (body) => {
      // TODO: users are not implemented yet
      if (1 > 0) {
        return {
          createdAt: new Date().toISOString(),
          id: 1,
          name: "test",
          updatedAt: new Date().toISOString(),
        };
      }

      const { data } = await axios<{ user: APIUser }>({
        method: "POST",
        url: `${BASE_URL}/auth/login`,
        withCredentials: true,
        data: body,
      });

      return data.user;
    },
  });
};
