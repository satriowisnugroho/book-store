CREATE TABLE "books" (
  "id" serial PRIMARY KEY,
  "isbn" varchar NOT NULL,
  "title" varchar NOT NULL,
  "price" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "books" ("title");
