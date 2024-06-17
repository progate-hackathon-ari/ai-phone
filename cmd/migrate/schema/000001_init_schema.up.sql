CREATE TABLE "users" (
  "user_id" varchar(63) PRIMARY KEY,
  "first_name" varchar(31),
  "last_name" varchar(31),
  "first_name_reading" varchar(31),
  "last_name_reading" varchar(31),
  "email" varchar(255) NOT NULL,
  "icon_path" varchar(255),
  "age" int,
  "gender" int,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "is_delete" boolean NOT NULL DEFAULT false
);

CREATE TABLE "user_roles" (
  "user_id" varchar(63) NOT NULL,
  "role_id" varchar(63) NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "roles" (
  "role_id" varchar(63) PRIMARY KEY,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "permissions" (
  "permission_id" varchar(63) PRIMARY KEY,
  "uri" text NOT NULL,
  "req_method" varchar(63) NOT NULL,
  "effect" varchar(10) NOT NULL
);

CREATE TABLE "role_permissions" (
  "role_id" varchar(63) NOT NULL,
  "permission_id" varchar(63) NOT NULL
);

CREATE TABLE "session" (
  "session_id" varchar(63) PRIMARY KEY,
  "user_agent" text NOT NULL,
  "user_id" varchar(63) NOT NULL,
  "token" text NOT NULL,
  "expiration_time" int NOT NULL
);

CREATE TABLE "schedules" (
  "schedule_id" varchar(63) PRIMARY KEY,
  "theater_id" varchar(63) NOT NULL,
  "movie_id" varchar(63) NOT NULL,
  "start_date" timestamptz NOT NULL
);

CREATE TABLE "theaters" (
  "theater_id" varchar(63) PRIMARY KEY,
  "theater_size_id" varchar(63) NOT NULL,
  "name" varchar(63) NOT NULL
);

CREATE TABLE "theaters_sizes" (
  "theater_size_id" varchar(63) PRIMARY KEY,
  "name" varchar(63) NOT NULL,
  "capacity" int NOT NULL
);

CREATE TABLE "movies" (
  "movie_id" varchar(63) PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "director" varchar(63) NOT NULL,
  "summary" text NOT NULL,
  "thumbnail_path" varchar(255) NOT NULL,
  "link" text NOT NULL,
  "term" int NOT NULL,
  "release_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "is_delete" boolean NOT NULL DEFAULT false
);

CREATE TABLE "movie_images" (
  "movie_id" varchar(63),
  "file_path" varchar(255) NOT NULL,
  "order" serial NOT NULL
);

CREATE TABLE "theaters_seats" (
  "theater_seat_id" varchar(63) PRIMARY KEY,
  "user_id" varchar(63) NOT NULL,
  "schedule_id" varchar(63) NOT NULL,
  "seat_name" varchar(31) NOT NULL
);

CREATE TABLE "orders" (
  "order_id" varchar(63) PRIMARY KEY,
  "user_id" varchar(63) NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "orders_details" (
  "order_detail_id" varchar(63) PRIMARY KEY,
  "order_id" varchar(63) NOT NULL,
  "theaters_seats_id" varchar(63) NOT NULL,
  "price_type_id" varchar(63) NOT NULL
);

CREATE TABLE "price_types" (
  "price_type_id" varchar(63) PRIMARY KEY,
  "name" varchar(63) NOT NULL,
  "price" int NOT NULL
);

CREATE UNIQUE INDEX ON "theaters_seats" ("schedule_id", "seat_name");

COMMENT ON COLUMN "movies"."term" IS '上映時間';

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "schedules" ADD FOREIGN KEY ("theater_id") REFERENCES "theaters" ("theater_id");

ALTER TABLE "schedules" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("movie_id");

ALTER TABLE "theaters_seats" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "theaters_seats" ADD FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("schedule_id");

ALTER TABLE "theaters" ADD FOREIGN KEY ("theater_size_id") REFERENCES "theaters_sizes" ("theater_size_id");

ALTER TABLE "orders_details" ADD FOREIGN KEY ("price_type_id") REFERENCES "price_types" ("price_type_id");

ALTER TABLE "orders_details" ADD FOREIGN KEY ("theaters_seats_id") REFERENCES "theaters_seats" ("theater_seat_id");

ALTER TABLE "orders_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "movie_images" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("movie_id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("permission_id");
