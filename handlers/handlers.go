package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// The last handler should return to the caller

func GetBooks(c *fiber.Ctx) error {
	slog.Info("get all books")
	return c.JSON(fiber.Map{"book1": "bookName1"})
}
func GetBookById(c *fiber.Ctx) error {
	slog.Info("get the book by Id")
	bookId := c.Params("id")
	slog.Info("info", "bookId", bookId)
	return c.JSON(fiber.Map{"bookId": bookId})
}
func GetAuthorById(c *fiber.Ctx) error {
	slog.Info("get the author by Id")
	authorId := c.Params("id")
	if authorId == "" {
		return c.JSON(fiber.Map{"authorId": "id is empty"}) //c.SendString("id is empty")
	}
	slog.Info("info", "authorId", authorId)
	return c.JSON(fiber.Map{"authorId": authorId})
}

// --------------== Item
// GetItems is an authenticated endpoint that returns all items
func GetItems(c *fiber.Ctx) error {
	slog.Info("GetItems", "path", c.Path(), "method", c.Method())
	return c.SendStatus(http.StatusOK)
}

// unauthenticated endpoint that gets user info and tries to authenticate them
func Login(c *fiber.Ctx) error {
	type loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req loginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	slog.Info("Login", "username", req.Username, "password", req.Password)
	if req.Username != "admin" || req.Password != "admin" {
		return c.Status(http.StatusUnauthorized).SendString("invalid credentials")
	}

	c.Response().Header.Add("Authorization", "Bearer 1234567890")
	slog.Info("login request recieved", "path", c.Path(), "method", c.Method())
	return c.SendStatus(http.StatusOK)
}
