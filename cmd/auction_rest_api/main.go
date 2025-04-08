package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var products map[string]Product
var mapperDTO DTOMapper

func main() {
	products = make(map[string]Product)
	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", getProducts)
		r.Post("/", addProduct)
		r.Get("/{name}", getProduct)
		r.Delete("/{name}", deleteProduct)
		r.Post("/{name}/sell", sellProduct)
	})

	fmt.Println("Server started listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	filterRequested := r.URL.Query().Get("sold")
	var soldFilter func(Product) bool

	if filterRequested == "true" {
		soldFilter = func(p Product) bool { return p.IsSold }
	} else if filterRequested == "false" {
		soldFilter = func(p Product) bool { return !p.IsSold }
	} else if filterRequested == "" {
		soldFilter = func(_ Product) bool { return true }
	}

	prodSlice := make([]SimpleProductDTO, 0)
	for _, prod := range products {
		if soldFilter(prod) {
			prodSlice = append(prodSlice, mapperDTO.productToSimpleProductDTO(prod))
		}
	}

	jsonProducts, err := json.Marshal(prodSlice)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server couldn't retrieve list of products. Please try again later!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProducts)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var prod Product
	err := json.NewDecoder(r.Body).Decode(&prod)

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

	if prod.Price == nil {
		reqerrs += "Cannot add a product without a price.\n"
	}

	if prod.Price != nil && *prod.Price < 0 {
		reqerrs += "Cannot add a product with a negative price.\n"
	}

	if reqerrs != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(reqerrs))
		return
	}

	_, exists := products[prod.Name]

	if exists {
		w.WriteHeader(http.StatusConflict)
		w.Write(fmt.Appendf(nil, "Name \"%s\" is already taken by a different product.", prod.Name))
		return
	}

	products[prod.Name] = prod

	jsonProd, err := json.Marshal(prod)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("For some reason, the server couldn't display the created object, but it HAS been created. If you see this, congrats! It should never be possible."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonProd)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	_, exists := products[name]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	delete(products, name)

	w.WriteHeader(http.StatusNoContent)
}

func sellProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	prod, exists := products[name]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	prod.IsSold = true

	products[name] = prod

	jsonProd, err := json.Marshal(prod)

	w.WriteHeader(http.StatusOK)

	if err != nil {
		w.Write([]byte("For some reason, the server couldn't display the sold item, but it HAS been sold. If you see this, congrats! It should never be possible."))
		return
	}

	w.Write(jsonProd)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	prod, exists := products[name]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	jsonProd, err := json.Marshal(prod)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server couldn't display the requested product, though it does exist. This should never happen.."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProd)
}
