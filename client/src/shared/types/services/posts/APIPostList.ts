import { RatingEnum } from "shared/types/enums/rating-enum";

export interface APIPostList {
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
    pools: any;
    tags: any;
    relations: any;
  }>;
  count: number;
}
