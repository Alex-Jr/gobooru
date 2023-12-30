import { BASE_URL } from "services/BASE_URL";

export const filePathToUrl = (s: string) => `${BASE_URL}/static${s}`;
