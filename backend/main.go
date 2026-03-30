// backend/main.go - Servidor básico Go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    // Permitir conexiones desde SvelteKit
    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173",
        AllowMethods: "GET,POST,PUT,DELETE",
    }))

    // Ruta de prueba
    app.Get("/api/saludo", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "mensaje": "🚀 Backend con Go funcionando!",
            "status":  "ok",
        })
    })

    app.Listen(":3000")
}