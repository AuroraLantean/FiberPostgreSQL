package main

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// --------------== Item
func getItems(c *fiber.Ctx) error {
	slog.Info("get items")
	itemPath := c.Params("*")
	if itemPath == "" {
		return c.JSON(fiber.Map{"itemPath": "itemPath is empty"})
	}
	slog.Info("info", "itemPath", itemPath)
	subPaths := strings.Split(itemPath, "/")
	for _, subsubPath := range subPaths {
		slog.Info("\t-item subPath", "subPath", subsubPath)
	}
	return c.JSON(fiber.Map{"itemPath": itemPath})
}

func getItemById(c *fiber.Ctx) error {
	item := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "Widget 1 for ACME",
	}

	return c.JSON(item)
}

const GRID = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" //goroutine Id
var idx = 0

func getHandlerId() string {
	c := GRID[idx%26]
	idx++
	return fmt.Sprintf("GRID-%v-%c", idx, c) // GRID-1-A, GRID-2-B
}
func getHandler(c *fiber.Ctx) error {
	ccId := getHandlerId()
	name := c.Params("name")
	go func() {
		slog.Info("starting handler", "ccId", ccId, "name", name)
		t := time.After(10 * time.Second)
		for {
			select {
			case <-t:
				slog.Info("hander done", "ccId", ccId, "name", name)
				return
			default:
				slog.Info("still running", "ccId", ccId, "name", name) //Set fiber.config {Immutable : true} to fix this name value!
				time.Sleep(1 * time.Second)
			}
		}
	}()
	slog.Info("request received", "name", name)
	return c.JSON(fiber.Map{"message": "Welcome"})
	//return c.SendString("Hello, World!!!")
	//return nil //hahakkkk
}
