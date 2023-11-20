CREATE TABLE "posts" (
  "id" SERIAL NOT NULL,
  "rating" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "tag_ids" TEXT[] NOT NULL,
  "tag_count" INTEGER NOT NULL,
  "pool_count" INTEGER NOT NULL,
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

CREATE TABLE "tags" (
  "id" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "post_count" INTEGER NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY ("id")
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