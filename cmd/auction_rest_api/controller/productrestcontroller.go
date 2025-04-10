package controller

import (
	"curs1_boilerplate/cmd/auction_rest_api/infrastructure"
	"curs1_boilerplate/cmd/auction_rest_api/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductRestController struct {
	dtoMapper   model.DTOMapper
	productRepo infrastructure.ProductRepository
}

func (rc *ProductRestController) SetupRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", rc.getProducts)
		r.Post("/", rc.addProduct)
		r.Get("/{name}", rc.getProduct)
		r.Delete("/{name}", rc.deleteProduct)
		r.Post("/{name}/sell", rc.sellProduct)
	})
}

func (rc *ProductRestController) getProducts(w http.ResponseWriter, r *http.Request) {
	filterRequested := r.URL.Query().Get("sold")
	var soldFilter func(model.Product) bool

	// TODO: Move this logic either in the repo or the service layer.
	if filterRequested == "true" {
		soldFilter = func(p model.Product) bool { return p.IsSold }
	} else if filterRequested == "false" {
		soldFilter = func(p model.Product) bool { return !p.IsSold }
	} else if filterRequested == "" {
		soldFilter = func(_ model.Product) bool { return true }
	}

	allProducts, err := rc.productRepo.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unknown server occured. Could not retrieve list of products."))
		return
	}

	prodSlice := make([]model.SimpleProductDTO, 0)
	for _, prod := range allProducts {
		if soldFilter(*prod) {
			prodSlice = append(prodSlice, rc.dtoMapper.ProductToSimpleProductDTO(*prod))
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

func (rc *ProductRestController) addProduct(w http.ResponseWriter, r *http.Request) {
	var prod model.Product
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

	addedProduct, err := rc.productRepo.Add(r.Context(), prod)

	if err != nil {
		// TODO: Don't just assume this means the name is taken. Make error for that specifically.
		w.WriteHeader(http.StatusConflict)
		w.Write(fmt.Appendf(nil, "Name \"%s\" is already taken by a different product.", prod.Name))
		return
	}

	jsonProd, err := json.Marshal(addedProduct)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("For some reason, the server couldn't display the created object, but it HAS been created. If you see this, congrats! It should never be possible."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonProd)
}

func (rc *ProductRestController) deleteProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	err := rc.productRepo.Delete(r.Context(), name)

	if err != nil {
		// TODO: Don't just assume this means no product was found with name. Make error specifically for this.
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rc *ProductRestController) sellProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	foundProduct, err := rc.productRepo.GetOne(r.Context(), name)
	if err != nil {
		// TODO: Don't assume that's the error. Make a specific error.
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	if foundProduct.IsSold {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("The product has already been sold."))
		return
	}

	foundProduct.IsSold = true
	updatedProduct, err := rc.productRepo.Update(r.Context(), *foundProduct)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	jsonProd, err := json.Marshal(updatedProduct)

	w.WriteHeader(http.StatusOK)

	if err != nil {
		w.Write([]byte("For some reason, the server couldn't display the sold item, but it HAS been sold. If you see this, congrats! It should never be possible."))
		return
	}

	w.Write(jsonProd)
}

func (rc *ProductRestController) getProduct(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	foundProduct, err := rc.productRepo.GetOne(r.Context(), name)
	if err != nil {
		// TODO: Don't assume that's the error. Make a specific error.
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	jsonProd, err := json.Marshal(foundProduct)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server couldn't display the requested product, though it does exist. This should never happen.."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProd)
}

func NewProductRestController(productRepository infrastructure.ProductRepository) *ProductRestController {
	return &ProductRestController{productRepo: productRepository}
}
