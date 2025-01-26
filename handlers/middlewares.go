package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/gofiber/fiber/v2"
)

// Middleware should return to the caller if error happens, or must call c.Next() to go to the next handler
func VerifyLogin(c *fiber.Ctx) error {
	// ensure the user is logged in
	token := c.Request().Header.Peek("Authorization")
	if token == nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	jwt := string(token)
	if jwt != "Bearer 1234567890" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid token"})
	}

	slog.Info("got request", "method", c.Method(), "path", c.Path(), "token", token)
	return c.Next()
}

func Logging(c *fiber.Ctx) error {
	// log request method, path, and param 'id'
	reqId := c.Request().Header.Peek("REQUEST-ID")
	slog.Info("got request", "method", c.Method(), "path", c.Path(), "requestId", reqId)
	return c.Next()
}

func RequestID(c *fiber.Ctx) error {
	// add a unique request ID to each request
	// https://github.com/Microsoft/go-winio
	uid, _ := guid.NewV4()
	c.Request().Header.Add("REQUEST-ID", uid.String())
	return c.Next()
}
