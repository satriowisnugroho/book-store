CREATE TABLE "users" (
  -- "id" bigint PRIMARY KEY AUTO_INCREMENT,
  "id" bigint PRIMARY KEY,
  "email" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "crypted_password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()) 
);

CREATE UNIQUE INDEX ON "users" ("email");
