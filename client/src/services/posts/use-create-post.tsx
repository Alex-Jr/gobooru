import { useMutation } from "@tanstack/react-query";
import axios from "axios";

import { BASE_URL } from "../BASE_URL";

interface IUseCreatePost {
  tags: string[];
  sources: string[];
  customs: string[];
  description: string;
  parentId?: string;
  rating: string;
  file: File | undefined;
}

interface IUseCreatePoolResponse {
  post: {
    id: number;
  };
}

export const useCreatePost = <T,>() => {
  return useMutation<IUseCreatePoolResponse, Error, IUseCreatePost>(
    ["edit", "post"],
    {
      mutationFn: async (data) => {
        const formData = new FormData();

        console.log(data.tags);

        data.tags.forEach((tag) => {
          formData.append("tags", tag);
        });

        data.sources.forEach((source) => {
          formData.append("sources", source);
        });

        if (data.description) {
          formData.append("description", data.description);
        }

        if (data.customs) {
          data.customs.forEach((c) => {
            formData.append("custom", c);
          });
        }

        if (data.parentId) {
          formData.append("parentId", data.parentId);
        }

        formData.append("rating", data.rating);
        formData.append("file", data.file as any);

        const { data: axiosData } = await axios({
          method: "POST",
          url: BASE_URL + "/posts",
          data: formData,
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });

        return axiosData;
      },
    }
  );
};
