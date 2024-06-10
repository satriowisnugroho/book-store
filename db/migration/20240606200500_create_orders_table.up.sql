CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "fee" integer NOT NULL,
  "total_price" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "orders" ("user_id", "created_at");
