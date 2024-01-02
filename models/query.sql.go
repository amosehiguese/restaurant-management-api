// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addOrderItems = `-- name: AddOrderItems :one
insert into orderItem (order_id, dish_id, quantity) values (
    $1, $2, $3
)
returning id, order_id, dish_id, quantity
`

type AddOrderItemsParams struct {
	OrderID  uuid.UUID `json:"order_id"`
	DishID   uuid.UUID `json:"dish_id"`
	Quantity int32     `json:"quantity"`
}

func (q *Queries) AddOrderItems(ctx context.Context, arg AddOrderItemsParams) (Orderitem, error) {
	row := q.db.QueryRowContext(ctx, addOrderItems, arg.OrderID, arg.DishID, arg.Quantity)
	var i Orderitem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.DishID,
		&i.Quantity,
	)
	return i, err
}

const cancelReservation = `-- name: CancelReservation :exec
delete from reservation
where id = $1
`

func (q *Queries) CancelReservation(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, cancelReservation, id)
	return err
}

const checkReservations = `-- name: CheckReservations :many
select id, table_id, reservation_date, reservation_time, status, created_at, updated_at from reservation
where status iLike '%available%'
`

func (q *Queries) CheckReservations(ctx context.Context) ([]Reservation, error) {
	rows, err := q.db.QueryContext(ctx, checkReservations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reservation
	for rows.Next() {
		var i Reservation
		if err := rows.Scan(
			&i.ID,
			&i.TableID,
			&i.ReservationDate,
			&i.ReservationTime,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const confirmReservation = `-- name: ConfirmReservation :exec
update reservation
set status = $2
where id = $1
`

type ConfirmReservationParams struct {
	ID     uuid.UUID         `json:"id"`
	Status ReservationStatus `json:"status"`
}

func (q *Queries) ConfirmReservation(ctx context.Context, arg ConfirmReservationParams) error {
	_, err := q.db.ExecContext(ctx, confirmReservation, arg.ID, arg.Status)
	return err
}

const createInvoice = `-- name: CreateInvoice :one
insert into invoice (order_id, invoice_date, total_amount, tax, discount, grand_total) values (
    $1, $2, $3, $4, $5, $6
)
returning id, order_id, invoice_date, total_amount, tax, discount, grand_total
`

type CreateInvoiceParams struct {
	OrderID     uuid.UUID `json:"order_id"`
	InvoiceDate time.Time `json:"invoice_date"`
	TotalAmount string    `json:"total_amount"`
	Tax         string    `json:"tax"`
	Discount    string    `json:"discount"`
	GrandTotal  string    `json:"grand_total"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice,
		arg.OrderID,
		arg.InvoiceDate,
		arg.TotalAmount,
		arg.Tax,
		arg.Discount,
		arg.GrandTotal,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.InvoiceDate,
		&i.TotalAmount,
		&i.Tax,
		&i.Discount,
		&i.GrandTotal,
	)
	return i, err
}

const createMenu = `-- name: CreateMenu :one
insert into menu (
    name, description
) values (
    $1, $2
)
returning id, name, description
`

type CreateMenuParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, createMenu, arg.Name, arg.Description)
	var i Menu
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createMenuDish = `-- name: CreateMenuDish :one
insert into dish (name, description, price, menu_id) values (
    $1, $2, $3, $4
)
returning id, name, description, price, menu_id
`

type CreateMenuDishParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	MenuID      uuid.UUID `json:"menu_id"`
}

func (q *Queries) CreateMenuDish(ctx context.Context, arg CreateMenuDishParams) (Dish, error) {
	row := q.db.QueryRowContext(ctx, createMenuDish,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.MenuID,
	)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.MenuID,
	)
	return i, err
}

const createOrder = `-- name: CreateOrder :one
insert into orders (status, created_at) values (
    $1, $2
)
returning id, status, created_at, updated_at
`

type CreateOrderParams struct {
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.Status, arg.CreatedAt)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createReservation = `-- name: CreateReservation :one
insert into reservation (table_id, reservation_date, reservation_time, status, created_at) values (
    $1, $2, $3, $4, $5
)
returning id, table_id, reservation_date, reservation_time, status, created_at, updated_at
`

type CreateReservationParams struct {
	TableID         uuid.UUID         `json:"table_id"`
	ReservationDate time.Time         `json:"reservation_date"`
	ReservationTime time.Time         `json:"reservation_time"`
	Status          ReservationStatus `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
}

func (q *Queries) CreateReservation(ctx context.Context, arg CreateReservationParams) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, createReservation,
		arg.TableID,
		arg.ReservationDate,
		arg.ReservationTime,
		arg.Status,
		arg.CreatedAt,
	)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.TableID,
		&i.ReservationDate,
		&i.ReservationTime,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createRole = `-- name: CreateRole :one
insert into roles (
    name, description
) values (
    $1, $2
)
returning id, name, description
`

type CreateRoleParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, arg.Name, arg.Description)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createTable = `-- name: CreateTable :one
insert into restaurant_table (name, capacity, status) values (
    $1, $2, $3
)
returning id, name, capacity, status
`

type CreateTableParams struct {
	Name     string                `json:"name"`
	Capacity int32                 `json:"capacity"`
	Status   RestaurantTableStatus `json:"status"`
}

func (q *Queries) CreateTable(ctx context.Context, arg CreateTableParams) (RestaurantTable, error) {
	row := q.db.QueryRowContext(ctx, createTable, arg.Name, arg.Capacity, arg.Status)
	var i RestaurantTable
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Capacity,
		&i.Status,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
insert into users (
    first_name, last_name, username, email, password_hash, user_role, created_at
) values (
    $1, $2, $3, $4, $5, $6, $7
)
returning id, first_name, last_name, username, email, password_hash, user_role, created_at, updated_at
`

type CreateUserParams struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	UserRole     int32     `json:"user_role"`
	CreatedAt    time.Time `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.UserRole,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteInvoice = `-- name: DeleteInvoice :exec
delete from invoice
where id = $1
`

func (q *Queries) DeleteInvoice(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteInvoice, id)
	return err
}

const deleteMenu = `-- name: DeleteMenu :exec
delete from menu
where id = $1
`

func (q *Queries) DeleteMenu(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteMenu, id)
	return err
}

const deleteMenuDish = `-- name: DeleteMenuDish :exec
delete from dish
where menu_id = $1 and id = $2
`

type DeleteMenuDishParams struct {
	MenuID uuid.UUID `json:"menu_id"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) DeleteMenuDish(ctx context.Context, arg DeleteMenuDishParams) error {
	_, err := q.db.ExecContext(ctx, deleteMenuDish, arg.MenuID, arg.ID)
	return err
}

const deleteOrder = `-- name: DeleteOrder :exec
delete from orders
where id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const deleteRole = `-- name: DeleteRole :exec
delete from roles
where id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRole, id)
	return err
}

const deleteTable = `-- name: DeleteTable :exec
delete from restaurant_table
where id = $1
`

func (q *Queries) DeleteTable(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTable, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users
where id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllInvoices = `-- name: GetAllInvoices :many
select id, order_id, invoice_date, total_amount, tax, discount, grand_total from invoice
`

func (q *Queries) GetAllInvoices(ctx context.Context) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, getAllInvoices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.InvoiceDate,
			&i.TotalAmount,
			&i.Tax,
			&i.Discount,
			&i.GrandTotal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMenu = `-- name: GetAllMenu :many
select id, name, description from menu
`

func (q *Queries) GetAllMenu(ctx context.Context) ([]Menu, error) {
	rows, err := q.db.QueryContext(ctx, getAllMenu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Menu
	for rows.Next() {
		var i Menu
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMenuDishes = `-- name: GetAllMenuDishes :many
select id, name, description, price, menu_id from dish
where menu_id = $1
order by price
`

func (q *Queries) GetAllMenuDishes(ctx context.Context, menuID uuid.UUID) ([]Dish, error) {
	rows, err := q.db.QueryContext(ctx, getAllMenuDishes, menuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dish
	for rows.Next() {
		var i Dish
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.MenuID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllOrderItems = `-- name: GetAllOrderItems :many
select id, order_id, dish_id, quantity from orderItem
where order_id = $1
`

func (q *Queries) GetAllOrderItems(ctx context.Context, orderID uuid.UUID) ([]Orderitem, error) {
	rows, err := q.db.QueryContext(ctx, getAllOrderItems, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Orderitem
	for rows.Next() {
		var i Orderitem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.DishID,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllOrders = `-- name: GetAllOrders :many
select id, status, created_at, updated_at from orders
`

func (q *Queries) GetAllOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getAllOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllReservations = `-- name: GetAllReservations :many
select id, table_id, reservation_date, reservation_time, status, created_at, updated_at from reservation
`

func (q *Queries) GetAllReservations(ctx context.Context) ([]Reservation, error) {
	rows, err := q.db.QueryContext(ctx, getAllReservations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reservation
	for rows.Next() {
		var i Reservation
		if err := rows.Scan(
			&i.ID,
			&i.TableID,
			&i.ReservationDate,
			&i.ReservationTime,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllRoles = `-- name: GetAllRoles :many
select id, name, description from roles
`

func (q *Queries) GetAllRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, getAllRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTables = `-- name: GetAllTables :many
select id, name, capacity, status from restaurant_table
`

func (q *Queries) GetAllTables(ctx context.Context) ([]RestaurantTable, error) {
	rows, err := q.db.QueryContext(ctx, getAllTables)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RestaurantTable
	for rows.Next() {
		var i RestaurantTable
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Capacity,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many
select id, first_name, last_name, username, email, password_hash, user_role, created_at, updated_at from users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.PasswordHash,
			&i.UserRole,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeSpecificOrderItem = `-- name: RemoveSpecificOrderItem :exec
delete from orderItem
where id = $1 and order_id = $2
`

type RemoveSpecificOrderItemParams struct {
	ID      uuid.UUID `json:"id"`
	OrderID uuid.UUID `json:"order_id"`
}

func (q *Queries) RemoveSpecificOrderItem(ctx context.Context, arg RemoveSpecificOrderItemParams) error {
	_, err := q.db.ExecContext(ctx, removeSpecificOrderItem, arg.ID, arg.OrderID)
	return err
}

const retrieveInvoice = `-- name: RetrieveInvoice :one
select id, order_id, invoice_date, total_amount, tax, discount, grand_total from invoice
where id = $1
`

func (q *Queries) RetrieveInvoice(ctx context.Context, id uuid.UUID) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, retrieveInvoice, id)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.InvoiceDate,
		&i.TotalAmount,
		&i.Tax,
		&i.Discount,
		&i.GrandTotal,
	)
	return i, err
}

const retrieveMenu = `-- name: RetrieveMenu :one
select id, name, description from menu
where id = $1 limit 1
`

func (q *Queries) RetrieveMenu(ctx context.Context, id uuid.UUID) (Menu, error) {
	row := q.db.QueryRowContext(ctx, retrieveMenu, id)
	var i Menu
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const retrieveMenuDish = `-- name: RetrieveMenuDish :one
select id, name, description, price, menu_id from dish 
where menu_id = $1 and id = $2
`

type RetrieveMenuDishParams struct {
	MenuID uuid.UUID `json:"menu_id"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) RetrieveMenuDish(ctx context.Context, arg RetrieveMenuDishParams) (Dish, error) {
	row := q.db.QueryRowContext(ctx, retrieveMenuDish, arg.MenuID, arg.ID)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.MenuID,
	)
	return i, err
}

const retrieveOrder = `-- name: RetrieveOrder :one
select id, status, created_at, updated_at from orders
where id = $1
`

func (q *Queries) RetrieveOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRowContext(ctx, retrieveOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const retrieveReservation = `-- name: RetrieveReservation :one
select id, table_id, reservation_date, reservation_time, status, created_at, updated_at from reservation
where id = $1
`

func (q *Queries) RetrieveReservation(ctx context.Context, id uuid.UUID) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, retrieveReservation, id)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.TableID,
		&i.ReservationDate,
		&i.ReservationTime,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const retrieveRole = `-- name: RetrieveRole :one
select id, name, description from roles
where id = $1 limit 1
`

func (q *Queries) RetrieveRole(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRowContext(ctx, retrieveRole, id)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const retrieveTable = `-- name: RetrieveTable :one
select id, name, capacity, status from restaurant_table
where id = $1
`

func (q *Queries) RetrieveTable(ctx context.Context, id uuid.UUID) (RestaurantTable, error) {
	row := q.db.QueryRowContext(ctx, retrieveTable, id)
	var i RestaurantTable
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Capacity,
		&i.Status,
	)
	return i, err
}

const retrieveUser = `-- name: RetrieveUser :one
select id, first_name, last_name, username, email, password_hash, user_role, created_at, updated_at from users
where id = $1 limit 1
`

func (q *Queries) RetrieveUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, retrieveUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const retrieveUserByEmail = `-- name: RetrieveUserByEmail :one
select id, first_name, last_name, username, email, password_hash, user_role, created_at, updated_at from users
where email = $1 limit 1
`

func (q *Queries) RetrieveUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, retrieveUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateInvoice = `-- name: UpdateInvoice :exec
update invoice
set invoice_date = $3,
total_amount = $4,
tax = $5,
discount = $6,
grand_total = $7
where id = $1 and order_id = $2
`

type UpdateInvoiceParams struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	InvoiceDate time.Time `json:"invoice_date"`
	TotalAmount string    `json:"total_amount"`
	Tax         string    `json:"tax"`
	Discount    string    `json:"discount"`
	GrandTotal  string    `json:"grand_total"`
}

func (q *Queries) UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) error {
	_, err := q.db.ExecContext(ctx, updateInvoice,
		arg.ID,
		arg.OrderID,
		arg.InvoiceDate,
		arg.TotalAmount,
		arg.Tax,
		arg.Discount,
		arg.GrandTotal,
	)
	return err
}

const updateMenu = `-- name: UpdateMenu :exec
update menu
set name = $2,
description = $3
where id = $1
`

type UpdateMenuParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (q *Queries) UpdateMenu(ctx context.Context, arg UpdateMenuParams) error {
	_, err := q.db.ExecContext(ctx, updateMenu, arg.ID, arg.Name, arg.Description)
	return err
}

const updateMenuDish = `-- name: UpdateMenuDish :exec
update dish 
set name = $3,
description = $4,
price = $5
where menu_id = $1 and id = $2
`

type UpdateMenuDishParams struct {
	MenuID      uuid.UUID `json:"menu_id"`
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
}

func (q *Queries) UpdateMenuDish(ctx context.Context, arg UpdateMenuDishParams) error {
	_, err := q.db.ExecContext(ctx, updateMenuDish,
		arg.MenuID,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
	)
	return err
}

const updateOrder = `-- name: UpdateOrder :exec
update orders
set status = $2
where id = $1
`

type UpdateOrderParams struct {
	ID     uuid.UUID   `json:"id"`
	Status OrderStatus `json:"status"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder, arg.ID, arg.Status)
	return err
}

const updateReservation = `-- name: UpdateReservation :exec
update reservation
set reservation_date = $3,
reservation_time = $4,
status = $5
where id = $1 and table_id = $2
`

type UpdateReservationParams struct {
	ID              uuid.UUID         `json:"id"`
	TableID         uuid.UUID         `json:"table_id"`
	ReservationDate time.Time         `json:"reservation_date"`
	ReservationTime time.Time         `json:"reservation_time"`
	Status          ReservationStatus `json:"status"`
}

func (q *Queries) UpdateReservation(ctx context.Context, arg UpdateReservationParams) error {
	_, err := q.db.ExecContext(ctx, updateReservation,
		arg.ID,
		arg.TableID,
		arg.ReservationDate,
		arg.ReservationTime,
		arg.Status,
	)
	return err
}

const updateRole = `-- name: UpdateRole :exec
update roles
set name = $2,
description = $3
where id = $1
`

type UpdateRoleParams struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateRole, arg.ID, arg.Name, arg.Description)
	return err
}

const updateTable = `-- name: UpdateTable :exec
update restaurant_table
set name = $2,
capacity = $3,
status = $4
where id = $1
`

type UpdateTableParams struct {
	ID       uuid.UUID             `json:"id"`
	Name     string                `json:"name"`
	Capacity int32                 `json:"capacity"`
	Status   RestaurantTableStatus `json:"status"`
}

func (q *Queries) UpdateTable(ctx context.Context, arg UpdateTableParams) error {
	_, err := q.db.ExecContext(ctx, updateTable,
		arg.ID,
		arg.Name,
		arg.Capacity,
		arg.Status,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
update users
set first_name = $2,
last_name = $3,
username = $4
where id = $1
`

type UpdateUserParams struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Username,
	)
	return err
}
