import axios from "axios";

import { BASE_URL } from "services/BASE_URL";
import { APICheck } from "shared/types/services/auth/APICheck";

export const authCheck = async (): Promise<APICheck["user"] | undefined> => {
  const { data } = await axios<APICheck>({
    method: "POST",
    url: `${BASE_URL}/auth/check`,
    withCredentials: true,
    validateStatus: (s) => s === 200 || s === 404,
  });

  if (!data) return undefined;

  return data.user;
};
