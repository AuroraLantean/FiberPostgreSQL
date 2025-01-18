package main

import (
	"log/slog"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

/*Fiber Framework: https://docs.gofiber.io/
https://github.com/gofiber/fiber

GORM, The fantastic ORM library: https://gorm.io/docs/
*/

func main() {
	slog.Info("main...")
	appConfig := fiber.Config{
		AppName:           "Golang PostgreSQL",
		EnablePrintRoutes: true,           //show routes talbe
		ServerHeader:      "Awesome App1", //so you know which server responded
		Immutable:         true,           //When set to true, this relinquishes the 0-allocation promise in certain cases in order to access the handler values (e.g. request bodies) in an immutable fashion so that these values are available even if you return from handler.Default: false
		//Prefork:           true,           //show child process IDs, so you know which process is doing what
		//CaseSesitive: true, //but do not do this
	}
	app := fiber.New(appConfig)
	app.Get("/special/:name", getHandler).Name("get default") // requires /:name
	app.Get("/books", getBooks).Name("get books")
	app.Get("/books/:id", getBookById).Name("get book by Id")
	app.Get("/authors/:id?", getAuthorById).Name("get author by Id")
	app.Get("/items/*", getItems).Name("get items with wildcard")

	log.Fatal(app.Listen(":3000"))
}
