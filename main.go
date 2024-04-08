package main

import (
	adapters "github.com/roku-on-it/golang-search/adapters/driven"
	"github.com/roku-on-it/golang-search/adapters/driving"
	"github.com/roku-on-it/golang-search/core/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	http := fiber.New()
	repo := adapters.NewUserRedisAdapter("localhost:6379")
	srv := services.NewUserService(repo)

	driving.SetupUserHTTPAdapter(srv, http.Group("/users"))
}
