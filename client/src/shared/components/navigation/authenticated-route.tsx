import { ReactNode, useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import { UserContext } from "shared/context/userContext";

import { BaseRoute } from "./base-route";

interface IBaseRoute {
  children: ReactNode;
}

export const AuthenticatedRoute = ({ children }: IBaseRoute) => {
  const { user } = useContext(UserContext);

  const navigate = useNavigate();

  useEffect(() => {
    if (!user) {
      navigate("/");
    }
  }, [user]);

  if (!user) return <></>;

  return <BaseRoute>{children}</BaseRoute>;
};
