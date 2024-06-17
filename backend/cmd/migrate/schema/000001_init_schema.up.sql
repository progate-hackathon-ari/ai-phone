CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "session" (
  "session_id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "token" text NOT NULL,
  "user_agent" text NOT NULL,
  "craeted_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL
);

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
