package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/alfuhigi/micro-order-api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func main() {
	if errr := v1(helpers.ConnectDB("DATABASE_URL", "silent")); errr != nil {
		log.Println(errr.Error())
	}
}

func v1(db *gorm.DB) error {

	config := fiber.Config{
		ServerHeader: "go",
	}

	for i := range os.Args[1:] {
		if os.Args[1:][i] == "-prefork" {
			config.Prefork = true
		}
	}

	app := *fiber.New(config)
	app.Use(logger.New())
	app.Use(recover.New())
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		for range interrupt {
			//TODO Notify Me
			app.Shutdown()
			os.Exit(1)
		}
	}()

	return app.Listen(":5800")
}
