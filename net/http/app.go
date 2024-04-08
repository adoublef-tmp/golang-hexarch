package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/roku-on-it/golang-search/internal/iam"
)

func App(iamDB iam.DB) *fiber.App {
	app := fiber.New()
	users := app.Group("/users")
	users.Get("/:id<guid>", handleUser(iamDB))
	users.Get("/:username<minLen(2);maxLen(32)>", handleSearchUser(iamDB))
	users.Post("/", handleJoinUser(iamDB))
	return app
}

func handleUser(iamDB iam.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uid, err := uuid.Parse(c.Params("id"))
		if err != nil { // 400 may not be the most correct code here
			return fiber.ErrBadRequest
		}

		found, err := iamDB.User(c.Context(), uid)
		if err != nil {
			if errors.Is(err, iam.ErrNotFound) {
				return fiber.ErrNotFound
			}
			return fiber.ErrInternalServerError
		}

		return c.JSON(found)
	}
}

func handleSearchUser(iamDB iam.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usr, err := iamDB.Search(c.Context(), c.Params("username"))
		if err != nil {
			if errors.Is(err, iam.ErrNotFound) {
				return fiber.ErrNotFound
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(usr)
	}
}

func handleJoinUser(iamDB iam.DB) fiber.Handler {
	type create struct {
		Display string `json:"display"`
		Name    string `json:"name"`
		Age     uint   `json:"age"`
	}

	type location struct {
		ID uuid.UUID `json:"id"`
	}
	return func(c *fiber.Ctx) error {
		var v create
		if err := c.BodyParser(&v); err != nil {
			return fiber.ErrBadRequest
		}

		uid, err := iamDB.OnBoard(c.Context(), v.Display, v.Name, v.Age)
		if err != nil {
			if errors.Is(err, iam.ErrDisplayExists) {
				return fiber.ErrConflict
			}

			return fiber.ErrInternalServerError
		}

		loc := &location{uid}
		return c.JSON(loc)
	}
}
