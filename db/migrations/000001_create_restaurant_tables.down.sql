DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS invoice;
DROP TABLE IF EXISTS reservation;
DROP TABLE IF EXISTS restaurant_table;
DROP TABLE IF EXISTS orderItem;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS dish;
DROP TABLE IF EXISTS menu;
DROP FUNCTION IF EXISTS update_updated_at();
DROP TRIGGER IF EXISTS e_u_ctd ON orders;
DROP TRIGGER IF EXISTS e_u_ctd ON reservation;
DROP TRIGGER IF EXISTS e_u_ctd ON users;
DROP TYPE IF EXISTS order_status;
DROP TYPE IF EXISTS reservation_status;
DROP TYPE IF EXISTS restaurant_table_status;



