CREATE TABLE "book" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "author" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "continent_name" varchar
);

CREATE TABLE "publisher" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "book_author" (
  "book_id" int,
  "author_id" int
);

CREATE TABLE "author_publisher" (
  "author_id" int,
  "publisher_id" int
);

CREATE TABLE "book_publisher" (
  "book_id" int,
  "publisher_id" int
);

ALTER TABLE "book_author" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "book_author" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");

ALTER TABLE "author_publisher" ADD FOREIGN KEY ("publisher_id") REFERENCES "publisher" ("id");

ALTER TABLE "author_publisher" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");

ALTER TABLE "book_publisher" ADD FOREIGN KEY ("book_id") REFERENCES "book" ("id");

ALTER TABLE "book_publisher" ADD FOREIGN KEY ("publisher_id") REFERENCES "publisher" ("id");
