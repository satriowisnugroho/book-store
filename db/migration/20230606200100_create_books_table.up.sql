CREATE TABLE "books" (
  -- "id" bigint PRIMARY KEY AUTO_INCREMENT,
  "id" bigint PRIMARY KEY,
  "isbn" varchar NOT NULL,
  "title" varchar NOT NULL,
  "price" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "books" ("isbn");
