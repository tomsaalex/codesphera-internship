package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	products := make([]Product, 0)
	router := chi.NewRouter()

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		jsonProducts, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Server couldn't retrieve list of products. Please try again later!"))
			return
		}

		w.Write([]byte(jsonProducts))
	})

	router.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		var prod Product
		err := json.NewDecoder(r.Body).Decode(&prod)
		fmt.Println(prod)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed request - body couldn't be parsed."))
			return
		}

		reqerrs := ""
		if prod.Name == "" {
			reqerrs += "Cannot add a product without a name.\n"
		}

		if prod.Description == "" {
			reqerrs += "Cannot add a product without a description.\n"
		}

		if prod.Price < 0 {
			reqerrs += "Cannot add a product with a negative price.\n"
		}

		if reqerrs != "" {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(reqerrs))
			return
		}

		products = append(products, prod)

		jsonProd, err := json.Marshal(prod)
		if err != nil {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("For some reason, the server couldn't display the object created, but it HAS been created. If you see this, congrats! It should never be possible."))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(jsonProd))
	})

	router.Delete("/")

	fmt.Println("Server started listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
