CREATE TABLE "posts" (
  "id" SERIAL NOT NULL,
  "rating" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "tag_ids" TEXT[] NOT NULL,
  "tag_count" INTEGER NOT NULL,
  "pool_count" INTEGER NOT NULL,
  "md5" TEXT NOT NULL,
  "file_ext" TEXT NOT NULL,
  "file_size" INTEGER NOT NULL,
  "file_path" TEXT NOT NULL,
  "thumb_path" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "pools" (
  "id" SERIAL NOT NULL,
  "name" TEXT NOT NULL,
  "post_count" INTEGER NOT NULL,
  "description" TEXT NOT NULL,
  "custom" TEXT[] NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("id")
);


CREATE TABLE "tag_categories" (
  "id" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "color" TEXT NOT NULL,
  "tag_count" INTEGER NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "tags" (
  "id" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "post_count" INTEGER NOT NULL,
  "category_id" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("category_id") REFERENCES "tag_categories" ("id") ON DELETE CASCADE
);

CREATE TABLE "pool_posts" (
  "pool_id" INTEGER NOT NULL,
  "post_id" INTEGER NOT NULL,
  "position" INTEGER NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("pool_id", "post_id"),
  FOREIGN KEY ("pool_id") REFERENCES "pools" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE
);

CREATE TABLE "post_tags" (
  "post_id" INTEGER NOT NULL,
  "tag_id" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("post_id", "tag_id"),
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE
);

CREATE TABLE "post_signatures" (
  "post_id" INTEGER NOT NULL,
  "signature" BYTEA NOT NULL,
  "words" TEXT[] NOT NULL,
  PRIMARY KEY ("post_id"),
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE
);

CREATE TABLE "post_relations" (
  "post_id" INTEGER NOT NULL,
  "other_post_id" INTEGER NOT NULL,
  "type" TEXT NOT NULL,
  "similarity" INTEGER NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("post_id", "other_post_id"),
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("other_post_id") REFERENCES "posts" ("id") ON DELETE CASCADE
);

CREATE TABLE "tag_aliases" (
  "tag_id" TEXT NOT NULL,
  "alias" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("tag_id", "alias"),
  FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE
);

CREATE TABLE "tag_implications" (
  "tag_id" TEXT NOT NULL,
  "implication_id" TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("tag_id", "implication_id"),
  FOREIGN KEY ("tag_id") REFERENCES "tags" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("implication_id") REFERENCES "tags" ("id") ON DELETE CASCADE
);
