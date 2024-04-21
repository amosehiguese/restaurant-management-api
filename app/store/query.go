package store

import (
	"context"
	"database/sql"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/google/uuid"
)

type Query interface {
	AddOrderItems(ctx context.Context, arg models.AddOrderItemsParams) (models.Orderitem, error)
	CancelReservation(ctx context.Context, id uuid.UUID) error
	CheckReservations(ctx context.Context) ([]models.Reservation, error)
	CreateInvoice(ctx context.Context, arg models.CreateInvoiceParams) (models.Invoice, error)
	CreateMenu(ctx context.Context, arg models.CreateMenuParams) (models.Menu, error)
	CreateMenuDish(ctx context.Context, arg models.CreateMenuDishParams) (models.Dish, error)
	CreateOrder(ctx context.Context, arg models.CreateOrderParams) (models.Order, error)
	CreateReservation(ctx context.Context, arg models.CreateReservationParams) (models.Reservation, error)
	CreateRole(ctx context.Context, arg models.CreateRoleParams) (models.Role, error)
	CreateTable(ctx context.Context, arg models.CreateTableParams) (models.RestaurantTable, error)
	CreateUser(ctx context.Context, arg models.CreateUserParams) (models.User, error)
	DeleteInvoice(ctx context.Context, id uuid.UUID) error
	DeleteMenu(ctx context.Context, id uuid.UUID) error
	DeleteMenuDish(ctx context.Context, arg models.DeleteMenuDishParams) error
	DeleteOrder(ctx context.Context, id uuid.UUID) error
	DeleteRole(ctx context.Context, id int32) error
	DeleteTable(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetAllInvoices(ctx context.Context) ([]models.Invoice, error)
	GetAllMenu(ctx context.Context) ([]models.Menu, error)
	GetAllMenuDishes(ctx context.Context, menuID uuid.UUID) ([]models.Dish, error)
	GetAllOrderItems(ctx context.Context, orderID uuid.UUID) ([]models.Orderitem, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	GetAllReservations(ctx context.Context) ([]models.Reservation, error)
	GetAllRoles(ctx context.Context) ([]models.Role, error)
	GetAllTables(ctx context.Context) ([]models.RestaurantTable, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	RemoveSpecificOrderItem(ctx context.Context, arg models.RemoveSpecificOrderItemParams) error
	RetrieveInvoice(ctx context.Context, id uuid.UUID) (models.Invoice, error)
	RetrieveMenu(ctx context.Context, id uuid.UUID) (models.Menu, error)
	RetrieveMenuDish(ctx context.Context, arg models.RetrieveMenuDishParams) (models.Dish, error)
	RetrieveOrder(ctx context.Context, id uuid.UUID) (models.Order, error)
	RetrieveReservation(ctx context.Context, id uuid.UUID) (models.Reservation, error)
	RetrieveRole(ctx context.Context, id int32) (models.Role, error)
	RetrieveTable(ctx context.Context, id uuid.UUID) (models.RestaurantTable, error)
	RetrieveUser(ctx context.Context, id uuid.UUID) (models.User, error)
	RetrieveUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateInvoice(ctx context.Context, arg models.UpdateInvoiceParams) error
	UpdateMenu(ctx context.Context, arg models.UpdateMenuParams) error
	UpdateMenuDish(ctx context.Context, arg models.UpdateMenuDishParams) error
	UpdateOrder(ctx context.Context, arg models.UpdateOrderParams) error
	UpdateOrderItem(ctx context.Context, arg models.UpdateOrderItemParams) error
	UpdateReservation(ctx context.Context, arg models.UpdateReservationParams) error
	UpdateRole(ctx context.Context, arg models.UpdateRoleParams) error
	UpdateTable(ctx context.Context, arg models.UpdateTableParams) error
	UpdateUser(ctx context.Context, arg models.UpdateUserParams) error
	WithTx(tx *sql.Tx) *models.Queries
}