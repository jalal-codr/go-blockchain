package routes

import (
	"fmt"
	"go-blockchain/controllers"
	"log"
	"net/http"
)

// InitRoutes sets up the routes and returns the router.
func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Define routes and their handlers
	mux.HandleFunc("/", controllers.HelloHandler)
	mux.HandleFunc("/createBlock", controllers.CreateBlock)

	mux.HandleFunc("/ws/mining", controllers.WebsocketConnection)

	mux.HandleFunc("/transferToken", controllers.TransferToken)

	mux.HandleFunc("/getBalance", controllers.GetBlockBalance)

	return mux
}

func StartServer() {
	router := InitRoutes()

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
