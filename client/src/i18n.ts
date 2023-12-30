import i18next from "i18next";
import { initReactI18next } from "react-i18next";

import { ptBr } from "locale/pt-br";

import { enUs } from "./locale/en-us";

i18next.use(initReactI18next).init({
  fallbackLng: "enUs",
  lng: "enUs",
  resources: {
    enUs: { translation: enUs },
    ptBr: { translation: ptBr },
  },
});

export default i18next;
