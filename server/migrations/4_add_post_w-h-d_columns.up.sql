-- using default because there's already data in the table
ALTER TABLE "posts" ADD COLUMN "width" INTEGER NOT NULL DEFAULT 0;
ALTER TABLE "posts" ADD COLUMN "height" INTEGER NOT NULL DEFAULT 0;
ALTER TABLE "posts" ADD COLUMN "duration" INTEGER NOT NULL DEFAULT 0;