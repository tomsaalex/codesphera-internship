package main

import (
	"curs1_boilerplate/controller"
	"curs1_boilerplate/db"
	"curs1_boilerplate/infrastructure"
	"curs1_boilerplate/middleware"
	"curs1_boilerplate/service"
	"curs1_boilerplate/util"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	// set up logger
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// connect to the db
	pool := db.NewConnectionPool()
	queries := db.New(pool)

	userRepository := infrastructure.NewDBUserRepository(queries)
	auctionRepository := infrastructure.NewDBAuctionRepository(queries)

	serviceDTOMapper := service.NewServiceDTOMapper()
	argonHelper := util.StandardArgon2idHash()
	userService := service.NewUserService(userRepository, *serviceDTOMapper, *argonHelper)
	auctionService := service.NewAuctionService(auctionRepository, userRepository, *serviceDTOMapper)

	jwtHelper := util.NewJwtUtil()
	userRestController := controller.NewUserRestController(*userService, *jwtHelper)

	auctionRestController := controller.NewAuctionRestController(*auctionService, *jwtHelper)

	generalRestController := controller.NewGeneralRestController(*jwtHelper)

	authRestController := controller.NewAuthRestController()

	r := chi.NewRouter()

	r.Use(middleware.AttachUser(*jwtHelper))
	r.Use(middleware.LoggingMiddleware)

	userRestController.SetupRoutes(r)
	generalRestController.SetupRoutes(r)
	authRestController.SetupRoutes(r)
	auctionRestController.SetupRoutes(r)

	// Serve images from ./static/images when hitting /images/*
	fs := http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images")))
	r.Handle("/images/*", fs)

	fmt.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
