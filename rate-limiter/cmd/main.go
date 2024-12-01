package main

import (
	"log"
	"path/filepath"
	limiter "rate-limiter/internal/limiter"
	middleware "rate-limiter/internal/middleware"

	fiber "github.com/gofiber/fiber/v2"
	viper "github.com/spf13/viper"
)

func init() {
	dir, _ := filepath.Abs("./")

	viper.AddConfigPath(dir)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
}

func main() {

	redisAddr := viper.GetString("REDIS_HOST") + ":" + viper.GetString("REDIS_PORT")
	storage := limiter.NewRedisStorage(redisAddr)

	app := fiber.New()
	app.Use(middleware.RateLimiterMiddleware(storage))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Println("Server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
