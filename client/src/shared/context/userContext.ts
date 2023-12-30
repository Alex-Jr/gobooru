import { Dispatch, SetStateAction, createContext } from "react";

import { APIUser } from "shared/types/services/user/APIUser";

export const UserContext = createContext<{
  user: APIUser | undefined;
  setUser: Dispatch<SetStateAction<APIUser | undefined>>;
}>({
  user: undefined,
  setUser: () => undefined,
});
