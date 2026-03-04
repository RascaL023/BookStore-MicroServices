package main

import (
	"fmt"
	"log"
	"net/http"
	"writer/internal/cache"
	"writer/internal/controller"
	"writer/internal/infrastructure/postgresql"
	"writer/internal/repository"
	"writer/internal/routes"
	"writer/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	uri := "postgres://rascal:@localhost:5555/mcr_writer"
	pool, err := postgresql.NewPostgrePool(uri)
	if err != nil { log.Fatal(err) }

	repo 		:= repository.New(pool)
	service 	:= service.New(repo)
	controller	:= controller.New(service)
	cache.InitRedis()

	r := chi.NewRouter()
	routes.RegisterRoutes(r, controller)

	fmt.Println("Server running on http://localhost:8085/writers")
	http.ListenAndServe(":8085", r)
}
