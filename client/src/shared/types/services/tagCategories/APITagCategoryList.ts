export interface ITagCategoryList {
  tag_categories: Array<{
    id: string;
    description: string;
    color: string;
    tag_count: number;
    created_at: string;
    updated_at: string;
  }>;
}
