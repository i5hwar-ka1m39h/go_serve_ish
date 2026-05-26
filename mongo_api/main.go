package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/config"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/controller"
	"github.com/i5hwar-ka1m39h/go_serve_ish/mongo_api/middleware"
)

// Main is the starting point of any executable Go program.
// In Go, the package must be named "main" and the starting function must be "main()".
func main() {
	// 1. Establish database connection and initialize MongoDB collections
	config.Connect()

	// 2. Instantiate Go's built-in HTTP Router (ServeMux)
	// mux acts as the central switchboard for the API. It matches the HTTP method
	// and URL path of incoming requests and forwards them to the correct controller function.
	mux := http.NewServeMux()

	// ============================================================================
	// PUBLIC ENDPOINTS (No authentication required)
	// ============================================================================
	
	// User Registration and Login
	mux.HandleFunc("POST /register", controller.RegisterHandler)
	mux.HandleFunc("POST /login", controller.LoginHandler)

	// Catalog Browsing
	mux.HandleFunc("GET /products", controller.ListProductsHandler)
	mux.HandleFunc("GET /products/{id}", controller.GetProductHandler)
	mux.HandleFunc("GET /categories", controller.ListCategoriesHandler)

	// ============================================================================
	// PROTECTED CUSTOMER ENDPOINTS (Require active authentication session)
	// ============================================================================
	
	// User Profile
	mux.Handle("GET /profile", middleware.Authenticate(http.HandlerFunc(controller.ProfileHandler)))

	// Shopping Cart
	mux.Handle("GET /cart", middleware.Authenticate(http.HandlerFunc(controller.GetCartHandler)))
	mux.Handle("POST /cart", middleware.Authenticate(http.HandlerFunc(controller.AddToCartHandler)))
	mux.Handle("DELETE /cart/{product_id}", middleware.Authenticate(http.HandlerFunc(controller.RemoveFromCartHandler)))

	// Order & Checkout
	mux.Handle("POST /orders", middleware.Authenticate(http.HandlerFunc(controller.PlaceOrderHandler)))
	mux.Handle("GET /orders", middleware.Authenticate(http.HandlerFunc(controller.ListOrdersHandler)))
	mux.Handle("GET /orders/{id}", middleware.Authenticate(http.HandlerFunc(controller.GetOrderHandler)))

	// ============================================================================
	// ADMINISTRATIVE ENDPOINTS (Protected: Requires Authenticated + Admin Role)
	// ============================================================================
	
	// Admin Product Management
	mux.Handle("POST /products", middleware.Authenticate(middleware.RequireAdmin(http.HandlerFunc(controller.CreateProductHandler))))
	mux.Handle("PUT /products/{id}", middleware.Authenticate(middleware.RequireAdmin(http.HandlerFunc(controller.UpdateProductHandler))))
	mux.Handle("DELETE /products/{id}", middleware.Authenticate(middleware.RequireAdmin(http.HandlerFunc(controller.DeleteProductHandler))))

	// Admin Category Management
	mux.Handle("POST /categories", middleware.Authenticate(middleware.RequireAdmin(http.HandlerFunc(controller.CreateCategoryHandler))))

	// ============================================================================
	// HTTP SERVER INITIALIZATION
	// ============================================================================
	
	// Wrap the entire router inside our Logger middleware.
	// This ensures that EVERY request gets logged to the console (GET, POST, error, latency, etc.)
	loggedMux := middleware.Logger(mux)

	// Retrieve port from .env config, default to "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("==================================================\n")
	fmt.Printf("🚀 E-COMMERCE BACKEND RUNNING ON PORT :%s\n", port)
	fmt.Printf("==================================================\n")

	// http.ListenAndServe takes a port address (e.g. ":8080") and our loggedMux handler.
	// It blocks execution and runs the server loop, listening forever for incoming requests.
	// If it fails (e.g. port is already in use), log.Fatal will print the error and exit the app.
	log.Fatal(http.ListenAndServe(":"+port, loggedMux))
}
