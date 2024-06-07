CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "book_id" bigint NOT NULL,
  "quantity" integer NOT NULL,
  "price" integer NOT NULL,
  "fee" integer NOT NULL,
  "total_price" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "orders" ("user_id");
