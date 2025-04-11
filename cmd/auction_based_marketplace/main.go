package main

import (
	"curs1_boilerplate/cmd/auction_based_marketplace/controller"
	"curs1_boilerplate/cmd/auction_based_marketplace/infrastructure"
	"curs1_boilerplate/cmd/auction_based_marketplace/service"
	"curs1_boilerplate/cmd/auction_based_marketplace/util"
	"curs1_boilerplate/db"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// connect to the db
	pool := db.NewConnectionPool()
	queries := db.New(pool)

	userRepository := infrastructure.NewDBUserRepository(queries)

	serviceDTOMapper := service.NewServiceDTOMapper()
	argonHelper := util.NewArgon2idHash(1, 32, 64*1024, 32, 256)
	userService := service.NewUserService(userRepository, *serviceDTOMapper, *argonHelper)

	userRestController := controller.NewUserRestController(*userService)

	r := chi.NewRouter()
	userRestController.SetupRoutes(r)

	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
