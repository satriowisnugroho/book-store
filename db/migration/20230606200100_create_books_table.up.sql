CREATE TABLE "books" (
  "id" serial PRIMARY KEY,
  "isbn" varchar NOT NULL,
  "title" varchar NOT NULL,
  "price" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "books" ("isbn");
