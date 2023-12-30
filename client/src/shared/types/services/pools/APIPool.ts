import { RatingEnum } from "shared/types/enums/rating-enum";

export interface APIPool {
  pool: {
    id: number;
    name: string;
    post_ids: Array<number>;
    post_count: number;
    description: string;
    custom: Array<any>;
    created_at: string;
    updated_at: string;
    posts: Array<{
      id: number;
      rating: RatingEnum;
      description: string;
      tag_ids: Array<string>;
      tag_count: number;
      pool_count: number;
      md5: string;
      file_ext: string;
      file_size: number;
      file_path: string;
      thumb_path: string;
      sources: Array<string>;
      custom: Array<string>;
      created_at: string;
      updated_at: string;
      pools: null;
      tags: null;
      relations: null;
    }>;
  };
}
