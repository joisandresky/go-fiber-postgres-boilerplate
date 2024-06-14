package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/goccy/go-json"

	"go-fiber-postgres-boilerplate/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *configs.Config
	logger *zap.SugaredLogger
	psqlDB *gorm.DB
}

func NewServer(
	cfg *configs.Config,
	logger *zap.SugaredLogger,
	psqlDB *gorm.DB,
) *Server {
	return &Server{
		cfg,
		logger,
		psqlDB,
	}
}

func (srv *Server) Run() error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(requestid.New())
	app.Use(fiberLog.New(fiberLog.Config{
		Format:     "[${pid} ${locals:requestid}] - [${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Singapore",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Welcome to %s", srv.cfg.App.Name),
		})
	})

	srv.InjectDependencies(app)

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})

	// 404 handler
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "404 Error: Lost in the digital wilderness. Maybe try a different map?",
		})
	})

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		port := fmt.Sprintf(":%v", srv.cfg.App.Port)

		go func() {
			srv.logger.Info(fmt.Sprintf("Starting HTTP server on port %s", port))
			if err := app.Listen(port); err != nil {
				srv.logger.Error("Error starting server: ", err)
			}
		}()

		// Wait for the termination signal or context cancellation
		<-ctx.Done()
		srv.logger.Info("Shutting down HTTP server...")
		if err := app.Shutdown(); err != nil {
			srv.logger.Error("Error shutting down HTTP server: ", err)
		}

		srv.logger.Info("HTTP Server Gracefully Shutdown!")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	srv.logger.Info("Received Signal, shutting down server and workers...")
	cancel()

	wg.Wait()

	close(sigChan)
	srv.logger.Info("Server Gracefully Shutdown!")
	return nil
}
