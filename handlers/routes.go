package handlers

import (
	"github.com/amosehiguese/restaurant-api/middleware"
	"github.com/go-chi/chi/v5"
)



func routes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/signup",SignUp)
		r.Post("/auth/login",SignIn)
		r.Get("/menu", GetMenu)
		

        // dishes
		r.Get("/menu/{id}/dishes", GetAllMenuDishes)
		r.Post("/menu/{id}/dishes", CreateMenuDish)
		r.Get("/menu/{id}/dishes/{dishID}", RetrieveMenuDish)
		r.Patch("/menu/{id}/dishes/{dishID}", UpdateMenuDish)
		r.Delete("/menu/{id}/dishes/{dishID}", DeleteMenuDish)

		// orders
		r.Get("/orders", GetAllOrders)
		r.Post("/menu/{id}/dishes/{dishID}/order", CreateOrder)
		r.Get("/orders/{id}", RetrieveOrder)
		r.Patch("/orders/{id}", UpdateOrder)
		r.Delete("/orders/{id}", DeleteOrder)

		r.Post("/orders/{id}/invoice", CreateInvoice)
		r.Patch("orders/{id}/invoices/{invoiceID}", UpdateInvoice)


		r.Get("/orders/{id}/items", GetAllOrderItems)
		r.Delete("/orders/{id}/items/{itemID}", RemoveSpecificOrderItem)
		
		// tables
		r.Get("/tables", GetAllTables)
		r.Post("/tables", CreateTable)
		r.Get("/tables/{id}", RetrieveTable)
		r.Patch("/tables/{id}", UpdateTable)
		r.Delete("/tables/{id}", DeleteTable)

		// reservations
		// r.Get("/reservations", GetAllReservations)
		// r.Post("/table/{id}/reservations/make", CreateReservation)
		// r.Get("/reservations/{id}", RetrieveReservation)
		// r.Patch("/table/{id}/reservations/{reservationID}", UpdateReservation)
		// r.Delete("/reservations/{id}/cancel", CancelReservation)

		// r.Post("/reservations/check-availability", CheckReservations)
		// r.Post("/reservations/{id}/confirm", ConfirmReservation)



		// invoices
		// r.Get("/invoices", GetAllInvoices)
		// r.Get("/invoices/{id}", RetrieveInvoice)
		// r.Delete("/invoices/{id}", DeleteInvoice)
		
		
		r.With(middleware.JWTAuthUser).Route("/p", func(r chi.Router) {
			
		},)

		r.With(middleware.JWTAuth).Route("/admin", func(r chi.Router) {
		

			r.Post("/menu", CreateMenu)
			r.Get("/menu/{id}", RetrieveMenu)
			r.Patch("/menu/{id}", UpdateMenu)
			r.Delete("/menu/{id}", DeleteMenu)
			// user
			r.Get("/users", GetUsers)
			r.Get("/users/{id}", RetrieveUser)
			r.Patch("/users/{id}", UpdateUser)
			// roles
			r.Get("/roles", GetRoles)
			r.Post("/roles", CreateRole)
			r.Patch("/role/{id}", UpdateRole)	
		},)
	})
}
