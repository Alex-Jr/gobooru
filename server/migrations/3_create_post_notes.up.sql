-- using default because there's already data in the table
CREATE TABLE "post_notes" (
  "id" SERIAL PRIMARY KEY,
  "post_id" INTEGER NOT NULL,
  "body" TEXT NOT NULL,
  "x" INTEGER NOT NULL,
  "y" INTEGER NOT NULL,
  "width" INTEGER NOT NULL,
  "height" INTEGER NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);