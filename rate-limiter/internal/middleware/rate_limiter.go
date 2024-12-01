package middleware

import (
	limiter "rate-limiter/internal/limiter"

	fiber "github.com/gofiber/fiber/v2"
)

func RateLimiterMiddleware(storage limiter.Storage) fiber.Handler {
	limiterByIP := &limiter.RateLimiter{Storage: storage, MaxRequests: 5, BlockDuration: 30}
	limiterByToken := &limiter.RateLimiter{Storage: storage, MaxRequests: 10, BlockDuration: 60}

	return func(c *fiber.Ctx) error {
		ip := c.IP()
		token := c.Get("Authorization")

		allowed, err := limiterByIP.Allow(ip)
		if err != nil {
			return c.Status(500).SendString("IP limiter error")
		}
		if !allowed {
			return c.Status(429).SendString("Too Many Requests")
		}

		allowed, err = limiterByToken.Allow(token)
		if err != nil {
			return c.Status(500).SendString("Token limiter error")
		}
		if !allowed {
			return c.Status(429).SendString("Too Many Requests")
		}

		return c.Next()
	}
}
