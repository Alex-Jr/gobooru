export interface APITagsList {
  tags: Array<{
    id: string;
    description: string;
    post_count: number;
    category_id: string;
    created_at: string;
    updated_at: string;
  }>;
  count: number;
}
