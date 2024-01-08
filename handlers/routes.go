package handlers

import (
	"net/http"

	"github.com/amosehiguese/restaurant-api/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)



func Routes(r *chi.Mux) {
	r.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})

	r.Get("/swagger*", httpSwagger.Handler())

	r.Route("/api/v1", func(r chi.Router) {
		// auth
		r.Post("/auth/signup",SignUp)
		r.Post("/auth/login",SignIn)

		// menu and dishes
		r.Get("/menu", GetMenu)
		r.Get("/menu/{id}", RetrieveMenu)
		r.Get("/menu/{id}/dishes", GetAllMenuDishes)
		r.Get("/menu/{id}/dishes/{dishID}", RetrieveMenuDish)

		// tables
		r.Get("/tables", GetAllTables)
		r.Get("/tables/{id}", RetrieveTable)
		
		
		r.With(middleware.JWTAuthUser).Route("/p", func(r chi.Router) {
			r.Post("/auth/signout",SignOut)
			r.Post("/token/renew", RenewTokens)

			// users
			r.Get("/users/{id}", RetrieveUser)

			// reservations
			r.Get("/reservations", GetAllReservations)
			r.Post("/reservations", CreateReservation)
			r.Get("/reservations/{id}", RetrieveReservation)
			r.Patch("/reservations/{id}", UpdateReservation)
			r.Delete("/reservations/{id}", DeleteReservation)			


			// invoices
			r.Post("/invoices", CreateInvoice)
			r.Get("/invoices/{id}", RetrieveInvoice)	
			
			// orders
			r.Get("/orders", GetAllOrders)
			r.Post("/orders", CreateOrder)
			r.Get("/orders/{id}", RetrieveOrder)
			r.Patch("/orders/{id}", UpdateOrder)
			r.Delete("/orders/{id}", DeleteOrder)
			
			// orderItems
			r.Get("/orders/{id}/items", GetAllOrderItems)
			r.Put("/orders/{id}/items", CreateOrderItem)
			r.Patch("/orders/{id}/items/{itemID}", UpdateOrderItem)
			r.Delete("/orders/{id}/items/{itemID}", RemoveSpecificOrderItem)							
		},)

		r.With(middleware.JWTAuth).Route("/admin", func(r chi.Router) {
		
			// menu
			r.Post("/menu", CreateMenu)
			r.Patch("/menu/{id}", UpdateMenu)
			r.Delete("/menu/{id}", DeleteMenu)


			// dishes
			r.Post("/menu/{id}/dishes", CreateMenuDish)
			r.Patch("/menu/{id}/dishes/{dishID}", UpdateMenuDish)
			r.Delete("/menu/{id}/dishes/{dishID}", DeleteMenuDish)		
			
			// tables
			r.Post("/tables", CreateTable)
			r.Patch("/tables/{id}", UpdateTable)
			r.Delete("/tables/{id}", DeleteTable)			

			// user
			r.Get("/users", GetUsers)
			r.Patch("/users/{id}", UpdateUser)

			// roles
			r.Get("/roles", GetRoles)
			r.Post("/roles", CreateRole)
			r.Patch("/roles/{id}", UpdateRole)	
			r.Delete("/roles/{id}", DeleteRole)	

			// invoices
			r.Get("/invoices", GetAllInvoices)	
			r.Patch("/invoices/{id}", UpdateInvoice)			
		},)
	})
}
