package main

import (
	"curs1_boilerplate/cmd/auction_based_marketplace/controller"
	"curs1_boilerplate/cmd/auction_based_marketplace/infrastructure"
	"curs1_boilerplate/cmd/auction_based_marketplace/service"
	"curs1_boilerplate/cmd/auction_based_marketplace/util"
	"curs1_boilerplate/db"
	"fmt"
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
	argonHelper := util.StandardArgon2idHash()
	userService := service.NewUserService(userRepository, *serviceDTOMapper, *argonHelper)

	jwtHelper := util.NewJwtUtil()
	userRestController := controller.NewUserRestController(*userService, *jwtHelper)

	r := chi.NewRouter()
	userRestController.SetupRoutes(r)

	fmt.Println("Server is listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
