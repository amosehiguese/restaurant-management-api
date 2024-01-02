-- name: GetAllUsers :many
select * from users;

-- name: RetrieveUser :one
select * from users
where id = $1 limit 1;

-- name: RetrieveUserByEmail :one
select * from users
where email = $1 limit 1;

-- name: UpdateUser :exec
update users
set first_name = $2,
last_name = $3,
username = $4
where id = $1;

-- name: DeleteUser :exec
delete from users
where id = $1;

-- name: CreateUser :one
insert into users (
    first_name, last_name, username, email, password_hash, user_role, created_at
) values (
    $1, $2, $3, $4, $5, $6, $7
)
returning *;

-- name: GetAllRoles :many
select * from roles;

-- name: RetrieveRole :one
select * from roles
where id = $1 limit 1;

-- name: CreateRole :one
insert into roles (
    name, description
) values (
    $1, $2
)
returning *;

-- name: UpdateRole :exec
update roles
set name = $2,
description = $3
where id = $1;

-- name: DeleteRole :exec
delete from roles
where id = $1;


-- name: GetAllMenu :many
select * from menu;

-- name: RetrieveMenu :one
select * from menu
where id = $1 limit 1;

-- name: CreateMenu :one
insert into menu (
    name, description
) values (
    $1, $2
)
returning *;

-- name: UpdateMenu :exec
update menu
set name = $2,
description = $3
where id = $1;

-- name: DeleteMenu :exec
delete from menu
where id = $1;


-- name: GetAllMenuDishes :many
select * from dish
where menu_id = $1
order by price;

-- name: CreateMenuDish :one
insert into dish (name, description, price, menu_id) values (
    $1, $2, $3, $4
)
returning *;

-- name: RetrieveMenuDish :one
select * from dish 
where menu_id = $1 and id = $2;

-- name: UpdateMenuDish :exec
update dish 
set name = $3,
description = $4,
price = $5
where menu_id = $1 and id = $2;

-- name: DeleteMenuDish :exec
delete from dish
where menu_id = $1 and id = $2;

-- name: AddOrderItems :one
insert into orderItem (order_id, dish_id, quantity) values (
    $1, $2, $3
)
returning *;

-- name: UpdateOrderItem :exec
update orderItem
set quantity = $4
where id = $1 and order_id = $2 and dish_id = $3;

-- name: GetAllOrderItems :many
select * from orderItem
where order_id = $1;


-- name: RemoveSpecificOrderItem :exec
delete from orderItem
where id = $1 and order_id = $2;

-- name: GetAllOrders :many
select * from orders;

-- name: CreateOrder :one
insert into orders (status, created_at) values (
    $1, $2
)
returning *;

-- name: RetrieveOrder :one
select * from orders
where id = $1;

-- name: UpdateOrder :exec
update orders
set status = $2
where id = $1;

-- name: DeleteOrder :exec
delete from orders
where id = $1;

-- name: GetAllTables :many
select * from restaurant_table;

-- name: CreateTable :one
insert into restaurant_table (name, capacity, status) values (
    $1, $2, $3
)
returning *;

-- name: RetrieveTable :one
select * from restaurant_table
where id = $1;

-- name: UpdateTable :exec
update restaurant_table
set name = $2,
capacity = $3,
status = $4
where id = $1;

-- name: DeleteTable :exec
delete from restaurant_table
where id = $1;

-- name: GetAllInvoices :many
select * from invoice;

-- name: RetrieveInvoice :one
select * from invoice
where id = $1;

-- name: CreateInvoice :one
insert into invoice (order_id, invoice_date, total_amount, tax, discount, grand_total) values (
    $1, $2, $3, $4, $5, $6
)
returning *;

-- name: UpdateInvoice :exec
update invoice
set total_amount = $3,
tax = $4,
discount = $5,
grand_total = $6
where id = $1 and order_id = $2;

-- name: DeleteInvoice :exec
delete from invoice
where id = $1;

-- name: GetAllReservations :many
select * from reservation;

-- name: CreateReservation :one
insert into reservation (table_id, reservation_date, reservation_time, status, created_at) values (
    $1, $2, $3, $4, $5
)
returning *;

-- name: RetrieveReservation :one
select * from reservation
where id = $1;

-- name: UpdateReservation :exec
update reservation
set reservation_date = $2,
reservation_time = $3,
status = $4
where id = $1;

-- name: CancelReservation :exec
delete from reservation
where id = $1;

-- name: CheckReservations :many
select * from reservation
where status iLike '%available%';

-- name: ConfirmReservation :exec
update reservation
set status = $2
where id = $1;
