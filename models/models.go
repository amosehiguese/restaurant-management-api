// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package models

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCanceled   OrderStatus = "canceled"
)

func (e *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatus(s)
	case string:
		*e = OrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatus: %T", src)
	}
	return nil
}

type NullOrderStatus struct {
	OrderStatus OrderStatus `json:"order_status"`
	Valid       bool        `json:"valid"` // Valid is true if OrderStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatus), nil
}

type ReservationStatus string

const (
	ReservationStatusAvailable ReservationStatus = "available"
	ReservationStatusConfirmed ReservationStatus = "confirmed"
	ReservationStatusCanceled  ReservationStatus = "canceled"
)

func (e *ReservationStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ReservationStatus(s)
	case string:
		*e = ReservationStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ReservationStatus: %T", src)
	}
	return nil
}

type NullReservationStatus struct {
	ReservationStatus ReservationStatus `json:"reservation_status"`
	Valid             bool              `json:"valid"` // Valid is true if ReservationStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullReservationStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ReservationStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ReservationStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullReservationStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ReservationStatus), nil
}

type RestaurantTableStatus string

const (
	RestaurantTableStatusOccupied  RestaurantTableStatus = "occupied"
	RestaurantTableStatusAvailable RestaurantTableStatus = "available"
)

func (e *RestaurantTableStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RestaurantTableStatus(s)
	case string:
		*e = RestaurantTableStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for RestaurantTableStatus: %T", src)
	}
	return nil
}

type NullRestaurantTableStatus struct {
	RestaurantTableStatus RestaurantTableStatus `json:"restaurant_table_status"`
	Valid                 bool                  `json:"valid"` // Valid is true if RestaurantTableStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRestaurantTableStatus) Scan(value interface{}) error {
	if value == nil {
		ns.RestaurantTableStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RestaurantTableStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRestaurantTableStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RestaurantTableStatus), nil
}

type Dish struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	MenuID      uuid.UUID `json:"menu_id"`
}

type Invoice struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	InvoiceDate time.Time `json:"invoice_date"`
	TotalAmount string    `json:"total_amount"`
	Tax         string    `json:"tax"`
	Discount    string    `json:"discount"`
	GrandTotal  string    `json:"grand_total"`
}

type Menu struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Order struct {
	ID        uuid.UUID    `json:"id"`
	Status    OrderStatus  `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Orderitem struct {
	ID       uuid.UUID `json:"id"`
	OrderID  uuid.UUID `json:"order_id"`
	DishID   uuid.UUID `json:"dish_id"`
	Quantity int32     `json:"quantity"`
}

type Reservation struct {
	ID              uuid.UUID         `json:"id"`
	TableID         uuid.UUID         `json:"table_id"`
	ReservationDate time.Time         `json:"reservation_date"`
	ReservationTime time.Time         `json:"reservation_time"`
	Status          ReservationStatus `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       sql.NullTime      `json:"updated_at"`
}

type RestaurantTable struct {
	ID       uuid.UUID             `json:"id"`
	Number   int32                 `json:"number"`
	Capacity int32                 `json:"capacity"`
	Status   RestaurantTableStatus `json:"status"`
}

type Role struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID           uuid.UUID    `json:"id"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"password_hash"`
	UserRole     int32        `json:"user_role"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}
