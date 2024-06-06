CREATE TABLE "books" (
  "id" bigint PRIMARY KEY,
  "isbn" varchar NOT NULL,
  "title" varchar NOT NULL,
  "price" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "books" ("isbn");
