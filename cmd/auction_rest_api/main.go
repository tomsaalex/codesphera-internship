package main

import (
	"curs1_boilerplate/cmd/auction_rest_api/controller"
	"curs1_boilerplate/cmd/auction_rest_api/infrastructure"
	"curs1_boilerplate/db"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//var mapperDTO model.DTOMapper

// Yes, I'm using the name as a primary key. Yes, I know it's inefficient. Yes, I know I could just make the name unique.
// I just don't see the benefit for this example. We're not doing anything with an ID.
func main() {
	// connect to the db
	pool := db.NewConnectionPool()
	queries := db.New(pool)

	productRepository := infrastructure.NewDBProductRepository(queries)
	//productRepository := infrastructure.NewMemoryProductRepository()
	productRestController := controller.NewProductRestController(productRepository)

	r := chi.NewRouter()
	productRestController.SetupRoutes(r)

	fmt.Println("Server started listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
