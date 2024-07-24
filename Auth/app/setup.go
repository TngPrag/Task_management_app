package app

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	//_ "user_manager/docs"
	"tele_auth/fs"
	"tele_auth/routers"

	//	event_processor "user_manager/logic/event_processor"
	_ "tele_auth/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupANDRun() error {
	fs.Fs_open()

	// Create app
	app := fiber.New()

	// Attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","), // Allow only specified HTTP methods
		AllowHeaders:     "Content-Type", // Allow only "Content-Type" header
		AllowCredentials: false,
	}))

	// Monitor the application
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Tele AuthZ Service"}))

	// Setup routes
	routers.SetupRoutes(app)

	// Safe shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	if err := app.Listen("localhost:8980"); err != nil {
		log.Panic(err)
	}

	serverShutdown.Wait()

	log.Println("Running cleanup tasks...")

	return nil
}
