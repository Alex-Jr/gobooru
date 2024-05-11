CREATE INDEX IF NOT EXISTS post_tags_post_id_idx ON public.post_tags USING btree (post_id);
CREATE INDEX IF NOT EXISTS post_tags_tag_id_idx ON public.post_tags USING btree (tag_id);
CREATE INDEX IF NOT EXISTS pool_posts_pool_id_idx ON public.pool_posts USING btree (pool_id);
CREATE INDEX IF NOT EXISTS pool_posts_post_id_idx ON public.pool_posts USING btree (post_id);
CREATE INDEX IF NOT EXISTS post_notes_post_id_idx ON public.post_notes USING btree (post_id);