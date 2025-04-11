package controller

import (
	"curs1_boilerplate/cmd/auction_based_marketplace/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserRestController struct {
	userService service.UserService
}

func NewUserRestController(userService service.UserService) *UserRestController {
	return &UserRestController{userService: userService}
}

func (rc *UserRestController) SetupRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", rc.registerUser)
		r.Post("/login", rc.loginUser)
	})
	r.Route("/user", func(r chi.Router) {
		r.Get("/", rc.getUser)
		r.Delete("/", rc.deleteUser)
		r.Put("/", rc.editUser)
	})
}

func (rc *UserRestController) registerUser(w http.ResponseWriter, r *http.Request) {
	var userDTO service.UserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed request - request body couldn't be parsed."))
		return
	}

	reqerrs := ""

	if userDTO.Fullname == "" {
		reqerrs += "Cannot register a user with no name.\n"
	}

	if userDTO.Email == "" {
		reqerrs += "Cannot register a user with no email.\n"
	}

	if userDTO.Password == "" {
		reqerrs += "Cannot register a user with no password.\n"
	}

	if reqerrs != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(reqerrs))
		return
	}

	err = rc.userService.Register(r.Context(), userDTO)
	if err != nil {
		// TODO: Don't just assume this means the name is taken. Make error for that specifically.
		// TODO: Also, this could mean email taken too.
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (rc *UserRestController) loginUser(w http.ResponseWriter, r *http.Request) {
	var userDTO service.UserLoginDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed request - request body couldn't be parsed."))
		return
	}

	reqerrs := ""

	if userDTO.Email == "" {
		reqerrs += "Users can't have a blank email.\n"
	}

	if userDTO.Password == "" {
		reqerrs += "Users can't have a blank password.\n"
	}

	if reqerrs != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(reqerrs))
		return
	}

	// TODO: This also kinda needs to do some jwt wizardry/cookie wizardry, but not quite yet
	err = rc.userService.Login(r.Context(), userDTO)

	if err != nil {
		// TODO: Don't just assume this means the name is taken. Make error for that specifically.
		// TODO: Also, this could mean email taken too.
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rc *UserRestController) deleteUser(w http.ResponseWriter, r *http.Request) {
	/*name := chi.URLParam(r, "name")

	err := rc.productRepo.Delete(r.Context(), name)

	if err != nil {
		// TODO: Don't just assume this means no product was found with name. Make error specifically for this.
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "There is no product with the name \"%s\".", name))
		return
	}

	w.WriteHeader(http.StatusNoContent)*/
}

func (rc *UserRestController) editUser(w http.ResponseWriter, r *http.Request) {
	/*name := chi.URLParam(r, "name")

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

	w.Write(jsonProd)*/
}

func (rc *UserRestController) getUser(w http.ResponseWriter, r *http.Request) {
	/*name := chi.URLParam(r, "name")

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
	w.Write(jsonProd)*/
}
