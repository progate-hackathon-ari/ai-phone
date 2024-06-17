
-- 外部キー制約を削除します
ALTER TABLE "session" DROP CONSTRAINT IF EXISTS "session_user_id_fkey";
ALTER TABLE "orders" DROP CONSTRAINT IF EXISTS "orders_user_id_fkey";
ALTER TABLE "schedules" DROP CONSTRAINT IF EXISTS "schedules_theater_id_fkey";
ALTER TABLE "schedules" DROP CONSTRAINT IF EXISTS "schedules_movie_id_fkey";
ALTER TABLE "theaters_seats" DROP CONSTRAINT IF EXISTS "theaters_seats_user_id_fkey";
ALTER TABLE "theaters_seats" DROP CONSTRAINT IF EXISTS "theaters_seats_schedule_id_fkey";
ALTER TABLE "theaters" DROP CONSTRAINT IF EXISTS "theaters_theater_size_id_fkey";
ALTER TABLE "orders_details" DROP CONSTRAINT IF EXISTS "orders_details_price_type_id_fkey";
ALTER TABLE "orders_details" DROP CONSTRAINT IF EXISTS "orders_details_theaters_seats_id_fkey";
ALTER TABLE "orders_details" DROP CONSTRAINT IF EXISTS "orders_details_order_id_fkey";
ALTER TABLE "movie_images" DROP CONSTRAINT IF EXISTS "movie_images_movie_id_fkey";

ALTER TABLE "user_roles" DROP CONSTRAINT IF EXISTS "user_roles_role_id_fkey";
ALTER TABLE "user_roles" DROP CONSTRAINT IF EXISTS "user_roles_user_id_fkey";
ALTER TABLE "role_permissions" DROP CONSTRAINT IF EXISTS "role_permissions_role_id_fkey";
ALTER TABLE "role_permissions" DROP CONSTRAINT IF EXISTS "role_permissions_permission_id_fkey";

-- テーブルを削除します
DROP TABLE IF EXISTS "session";
DROP TABLE IF EXISTS "user_roles";
DROP TABLE IF EXISTS "roles";
DROP TABLE IF EXISTS "permissions";
DROP TABLE IF EXISTS "role_permissions";
DROP TABLE IF EXISTS "orders_details";
DROP TABLE IF EXISTS "orders";
DROP TABLE IF EXISTS "theaters_seats";
DROP TABLE IF EXISTS "movie_images";
DROP TABLE IF EXISTS "movies";
DROP TABLE IF EXISTS "schedules";
DROP TABLE IF EXISTS "theaters";
DROP TABLE IF EXISTS "theaters_sizes";
DROP TABLE IF EXISTS "price_types";
DROP TABLE IF EXISTS "users";