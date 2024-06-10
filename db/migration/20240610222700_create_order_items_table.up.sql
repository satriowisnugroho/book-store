CREATE TABLE "order_items" (
  "id" serial PRIMARY KEY,
  "order_id" integer NOT NULL,
  "book_id" integer NOT NULL,
  "quantity" integer NOT NULL,
  "price" integer NOT NULL,
  "total_item_price" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "order_items" ("order_id");
