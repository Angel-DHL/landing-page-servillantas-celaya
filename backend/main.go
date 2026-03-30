// backend/main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ============ MODELOS ============

type ContactForm struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
	Servicio string `json:"servicio"`
	Mensaje  string `json:"mensaje"`
}

type Servicio struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Icono       string `json:"icono"`
}

type BlogPost struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Resumen   string `json:"resumen"`
	Imagen    string `json:"imagen"`
	Fecha     string `json:"fecha"`
	Categoria string `json:"categoria"`
	Slug      string `json:"slug"`
}

// ============ DATOS (después los migras a base de datos) ============

var servicios = []Servicio{
	{1, "Venta de Llantas", "Las mejores marcas nacionales e internacionales. Llantas para todo tipo de vehículo: automóvil, camioneta, SUV y carga.", "🛞"},
	{2, "Alineación Computarizada", "Equipos de última generación para una alineación precisa. Mejora el manejo y prolonga la vida de tus llantas.", "🎯"},
	{3, "Balanceo", "Eliminamos vibraciones para un manejo suave y seguro. Balanceo dinámico con tecnología de punta.", "⚖️"},
	{4, "Frenos", "Revisión, mantenimiento y cambio de balatas, discos y sistema de frenado completo. Tu seguridad es primero.", "🛑"},
	{5, "Suspensión", "Diagnóstico y reparación de amortiguadores, rotulas, terminales y todo el sistema de suspensión.", "🔧"},
	{6, "Cambio de Aceite", "Aceites sintéticos y minerales de las mejores marcas. Protege el motor de tu vehículo.", "🛢️"},
}

var blogPosts = []BlogPost{
	{1, "¿Cada cuánto debo cambiar mis llantas?",
		"Descubre las señales que indican que es momento de cambiar tus llantas y cómo prolongar su vida útil.",
		"/images/blog/post1.jpg", "2024-12-15", "Mantenimiento", "cada-cuanto-cambiar-llantas"},
	{2, "Importancia de la alineación y balanceo",
		"Una mala alineación puede costarte más de lo que imaginas. Conoce por qué es vital este servicio.",
		"/images/blog/post2.jpg", "2024-12-01", "Seguridad", "importancia-alineacion-balanceo"},
	{3, "5 tips para cuidar tus llantas en temporada de lluvias",
		"La temporada de lluvias puede ser peligrosa. Estos consejos te ayudarán a mantener el control.",
		"/images/blog/post3.jpg", "2024-11-20", "Tips", "tips-llantas-lluvias"},
}

// ============ HANDLERS ============

func handleGetServicios(c *fiber.Ctx) error {
	return c.JSON(servicios)
}

func handleGetBlog(c *fiber.Ctx) error {
	return c.JSON(blogPosts)
}

func handleContacto(c *fiber.Ctx) error {
	form := new(ContactForm)

	if err := c.BodyParser(form); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Datos inválidos. Verifica la información.",
		})
	}

	// Validaciones
	if form.Nombre == "" || form.Telefono == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "El nombre y teléfono son obligatorios.",
		})
	}

	// Aquí puedes guardar en BD o enviar email
	// Por ahora solo logueamos el lead
	log.Printf("📩 NUEVO LEAD [%s]", time.Now().Format("2006-01-02 15:04:05"))
	log.Printf("   Nombre:   %s", form.Nombre)
	log.Printf("   Email:    %s", form.Email)
	log.Printf("   Teléfono: %s", form.Telefono)
	log.Printf("   Servicio: %s", form.Servicio)
	log.Printf("   Mensaje:  %s", form.Mensaje)
	fmt.Println("─────────────────────────────────────")

	return c.JSON(fiber.Map{
		"success": true,
		"message": "¡Gracias por contactarnos! Te llamaremos pronto.",
	})
}

// ============ MAIN ============

func main() {
	app := fiber.New(fiber.Config{
		AppName: "ServiLlantas Celaya API v1.0",
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:4173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Rutas API
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "🚗 ServiLlantas API funcionando correctamente",
		})
	})

	api.Get("/servicios", handleGetServicios)
	api.Get("/blog", handleGetBlog)
	api.Post("/contacto", handleContacto)

	// Iniciar servidor
	fmt.Println("🚀 ServiLlantas API corriendo en http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
