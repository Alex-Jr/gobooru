import { useQuery } from "@tanstack/react-query";
import axios from "axios";

import { ITagCategoryList } from "shared/types/services/tagCategories/APITagCategoryList";

import { BASE_URL } from "../BASE_URL";

export const useTagCategoryList = () => {
  const { data } = useQuery({
    queryKey: ["tags-category"],
    queryFn: async () => {
      const { data } = await axios<ITagCategoryList>({
        method: "GET",
        url: `${BASE_URL}/tag-categories`,
      });

      return data;
    },
  });

  if (!data)
    return {
      tagCategories: [],
      tagCategoriesObj: {},
    };

  return {
    tagCategories: data.tag_categories,
    tagCategoriesObj: data.tag_categories.reduce((acc, cur) => {
      acc[cur.id] = cur.color;
      return acc;
    }, {} as Record<string, string>),
  };
};
