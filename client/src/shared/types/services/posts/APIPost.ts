export interface APIPost {
  post: {
    id: number;
    rating: string;
    status: string;
    description: string;
    custom: Array<string>;
    sources: Array<string>;
    tag_ids: Array<string>;
    tag_count: number;
    pool_count: number;
    md5: string;
    file_ext: string;
    file_size: number;
    file_path: string;
    thumb_path: string;
    created_at: string;
    updated_at: string;
    width: number;
    height: number;
    duration: number;
    pools: Array<{
      id: number;
      name: string;
      post_count: number;
      post_ids: Array<number>;
      description: string;
      custom: Array<string>;
      created_at: string;
      updated_at: string;
      posts: null;
    }>;
    tags: Array<{
      id: string;
      description: string;
      post_count: number;
      category_id: string;
      created_at: string;
      updated_at: string;
    }>;
    relations: Array<{
      created_at: string;
      other_post_id: number;
      other_post: {
        id: number;
        rating: string;
        description: string;
        tag_ids: Array<string>;
        tag_count: number;
        pool_count: number;
        md5: string;
        file_ext: string;
        file_size: number;
        file_path: string;
        thumb_path: string;
        created_at: string;
        updated_at: string;
        pools: any;
        tags: any;
        relations: any;
      };
      post_id: number;
      similarity: number;
      type: string;
    }>;
    notes: {
      id: number;
      post_id: number;
      body: string;
      x: number;
      y: number;
      width: number;
      height: number;
      created_at: string;
      updated_at: string;
    }[];
  };
}
