package controller

import (
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/service"
	"curs1_boilerplate/sharederrors"
	"curs1_boilerplate/util"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserRestController struct {
	userService service.UserService
	jwtHelper   util.JwtUtil
}

func NewUserRestController(userService service.UserService, jwtHelper util.JwtUtil) *UserRestController {
	return &UserRestController{
		userService: userService,
		jwtHelper:   jwtHelper,
	}
}

func (rc *UserRestController) SetupRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", rc.registerUser)
		r.Post("/login", rc.loginUser)
	})
	r.Route("/user", func(r chi.Router) {
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

	_, err = rc.userService.Register(r.Context(), userDTO)
	if err != nil {
		var authErr *service.AuthError
		var duplicateEntityErr *sharederrors.DuplicateEntityError
		var valErr *service.ValidationError

		if errors.As(err, &authErr) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The provided password does not respect requirements and cannot be used."))
		} else if errors.As(err, &duplicateEntityErr) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("There's already a user with the given email address. Please choose another one."))
		} else if errors.As(err, &valErr) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The provided user data is invalid and cannot be used"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An unexpected error occurred on our end. Please retry later!"))
		}
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

	// TODO: This also kinda needs to do some jwt wizardry/cookie wizardry, but not quite yet
	loggedUser, err := rc.userService.Login(r.Context(), userDTO)

	if err != nil {
		var authErr *service.AuthError
		var entityNotFoundErr *infrastructure.EntityNotFoundError
		var valErr *service.ValidationError

		if errors.As(err, &authErr) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authentication data invalid for the given email address."))
		} else if errors.As(err, &entityNotFoundErr) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No user found with the given email address."))
		} else if errors.As(err, &valErr) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The provided user data is invalid and cannot be used"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An unexpected error occurred on our end. Please retry later!"))
		}
		return
	}

	token, err := rc.jwtHelper.GenerateJWT(loggedUser.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to log user in."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
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
