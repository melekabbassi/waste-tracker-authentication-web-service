package main

import (
	"example/waste-tracker-authentication-web-service/database"
	"example/waste-tracker-authentication-web-service/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loadENV()

	//Initialize Fiber app
	app := generateApp()

	db := database.OpenDB()
	defer database.CloseDB(db)

	app.Listen(":8081")
}

func loadENV() error {
	goENV := os.Getenv("GO_ENV")
	if goENV == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func generateApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	// create healthcheck route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the dumptruck group and routes
	userGroup := app.Group("/users")
	userGroup.Get("/", handlers.GetUsers)
	userGroup.Get("/:id", handlers.GetUser)
	userGroup.Post("/", handlers.CreateUser)
	userGroup.Put("/:id", handlers.UpdateUser)
	userGroup.Delete("/:id", handlers.DeleteUser)

	return app
}
