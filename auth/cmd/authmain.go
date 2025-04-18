package main

import (
	"auth/internal/config"
	"auth/internal/database"
	"auth/internal/router"
	"context"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func main() {
	config.LoadEnvs()

	ctx := context.Background()
	database.Connect(ctx)
	defer database.ConnPool.Close()

	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":" + string(config.Env.Port))

	mu := sync.RWMutex{}
	mu.Rlo
}