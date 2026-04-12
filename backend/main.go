// backend/main.go
// Reemplaza TODO el archivo con esta versión actualizada:

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
	ID        int      `json:"id"`
	Titulo    string   `json:"titulo"`
	Resumen   string   `json:"resumen"`
	Contenido string   `json:"contenido"`
	Imagen    string   `json:"imagen"`
	Fecha     string   `json:"fecha"`
	Categoria string   `json:"categoria"`
	Slug      string   `json:"slug"`
	Autor     string   `json:"autor"`
	AutorRol  string   `json:"autor_rol"`
	Duracion  string   `json:"duracion"`
	Tags      []string `json:"tags"`
}

// ============ DATOS ============

var servicios = []Servicio{
	{1, "Venta de Llantas", "Las mejores marcas nacionales e internacionales.", "🛞"},
	{2, "Alineación Computarizada", "Equipos de última generación.", "🎯"},
	{3, "Balanceo", "Balanceo dinámico con tecnología de punta.", "⚖️"},
	{4, "Frenos", "Revisión y mantenimiento del sistema de frenado.", "🛑"},
	{5, "Suspensión", "Diagnóstico y reparación completa.", "🔧"},
	{6, "Cambio de Aceite", "Aceites sintéticos y minerales premium.", "🛢️"},
}

var blogPosts = []BlogPost{
	{
		ID:      1,
		Titulo:  "¿Cada cuánto debo cambiar mis llantas?",
		Resumen: "Descubre las señales que indican que es momento de cambiar tus llantas y cómo prolongar su vida útil con cuidados simples.",
		Contenido: `Las llantas son uno de los componentes más importantes de tu vehículo. Son el único punto de contacto entre tu auto y el camino, por lo que mantenerlas en buen estado es fundamental para tu seguridad.

## ¿Cuándo cambiar tus llantas?

**La regla general** es cambiar las llantas cada 40,000 a 60,000 kilómetros, dependiendo del tipo de llanta, tu estilo de manejo y las condiciones del camino.

Sin embargo, hay señales claras que indican que necesitas cambiarlas antes:

### 1. El indicador de desgaste
Todas las llantas modernas tienen indicadores de desgaste (TWI). Son pequeñas barras de goma ubicadas en los surcos de la llanta. Cuando la superficie de rodamiento está al mismo nivel que estos indicadores, es momento de cambiarlas.

### 2. La prueba de la moneda
Inserta una moneda de $1 peso en el surco de la llanta con el número hacia abajo. Si puedes ver el número completo, tus llantas necesitan reemplazo.

### 3. Grietas o abultamientos
Revisa las paredes laterales de tus llantas. Si notas grietas, cortes o abultamientos, cámbialas inmediatamente ya que podrían reventar.

### 4. Vibración excesiva
Si sientes vibraciones inusuales al manejar, podría ser señal de llantas desgastadas de forma irregular o dañadas internamente.

### 5. Antigüedad
Aunque no las uses mucho, las llantas se degradan con el tiempo. Se recomienda cambiarlas cada **5 años** como máximo, sin importar su apariencia.

## Consejos para prolongar la vida de tus llantas

- **Mantén la presión correcta**: Revisa la presión al menos una vez al mes
- **Realiza la rotación**: Cada 10,000 km rota las llantas de posición
- **Alineación y balanceo**: Hazlo cada 10,000 km o cuando notes desgaste irregular
- **Manejo suave**: Evita acelerones y frenazos bruscos
- **Evita baches**: Los impactos fuertes dañan la estructura interna

## ¿Necesitas llantas nuevas?

En **ServiLlantas Celaya** tenemos las mejores marcas a los mejores precios. Te asesoramos para elegir la llanta ideal según tu vehículo y estilo de manejo. ¡Visítanos!`,
		Imagen:    "/images/blog/llantas.jpg",
		Fecha:     "2024-12-15",
		Categoria: "Mantenimiento",
		Slug:      "cada-cuanto-cambiar-llantas",
		Autor:     "Carlos Ramírez",
		AutorRol:  "Director General",
		Duracion:  "5 min de lectura",
		Tags:      []string{"llantas", "mantenimiento", "seguridad", "consejos"},
	},
	{
		ID:      2,
		Titulo:  "Importancia de la alineación y balanceo",
		Resumen: "Una mala alineación puede costarte más de lo que imaginas. Conoce por qué es vital este servicio para tu seguridad y tu bolsillo.",
		Contenido: `La alineación y el balanceo son dos servicios que muchos conductores pasan por alto, pero que tienen un impacto enorme en la seguridad, el confort y la economía de tu vehículo.

## ¿Qué es la alineación?

La alineación consiste en ajustar los ángulos de las ruedas para que estén **perpendiculares al suelo y paralelas entre sí**. Esto garantiza que tu vehículo se desplace en línea recta y que las llantas hagan contacto uniforme con el pavimento.

### Señales de que necesitas alineación:

- El volante vibra o jala hacia un lado
- Desgaste disparejo en las llantas
- El volante no regresa al centro después de una curva
- Acabas de pasar por un bache fuerte

## ¿Qué es el balanceo?

El balanceo compensa las diferencias de peso en el conjunto llanta-rin. Cuando hay desbalanceo, se generan vibraciones que afectan la comodidad y aceleran el desgaste de componentes.

### Beneficios del balanceo correcto:

- Manejo suave sin vibraciones
- Mayor vida útil de las llantas
- Menor desgaste de amortiguadores
- Ahorro de combustible

## ¿Cada cuánto hacerlos?

| Servicio | Frecuencia recomendada |
|----------|----------------------|
| Alineación | Cada 10,000 km o 6 meses |
| Balanceo | Cada 10,000 km o al cambiar llantas |

## El costo de NO hacerlos

Una mala alineación puede reducir la vida útil de tus llantas hasta en un **50%**. Considerando que un juego de llantas cuesta varios miles de pesos, invertir en alineación y balanceo regulares te ahorra mucho dinero a largo plazo.

## Confía en los expertos

En **ServiLlantas Celaya** contamos con equipo de alineación 3D de última generación que garantiza precisión milimétrica. ¡Agenda tu cita hoy!`,
		Imagen:    "/images/blog/alineacion.jpg",
		Fecha:     "2024-12-01",
		Categoria: "Seguridad",
		Slug:      "importancia-alineacion-balanceo",
		Autor:     "Miguel Torres",
		AutorRol:  "Jefe de Taller",
		Duracion:  "4 min de lectura",
		Tags:      []string{"alineación", "balanceo", "seguridad", "ahorro"},
	},
	{
		ID:      3,
		Titulo:  "5 tips para cuidar tus llantas en temporada de lluvias",
		Resumen: "La temporada de lluvias puede ser peligrosa. Estos consejos te ayudarán a mantener el control de tu vehículo en condiciones adversas.",
		Contenido: `La temporada de lluvias presenta desafíos especiales para los conductores. El pavimento mojado reduce significativamente la adherencia de las llantas, aumentando el riesgo de accidentes. Aquí te compartimos 5 tips esenciales.

## 1. Verifica la profundidad del dibujo

El dibujo de la llanta es responsable de **evacuar el agua** entre la llanta y el pavimento. Con un dibujo desgastado, el agua no se canaliza correctamente y se produce el temido **aquaplaning** (la llanta "flota" sobre el agua).

**Profundidad mínima recomendada:** 3mm para condiciones de lluvia.

## 2. Mantén la presión correcta

Con lluvia, la presión correcta es aún más importante:

- **Presión baja**: Aumenta la superficie de contacto pero reduce la capacidad de evacuar agua
- **Presión alta**: Reduce la superficie de contacto y la adherencia
- **Presión correcta**: La que indica el fabricante de tu vehículo

Revisa la presión **en frío**, al menos una vez por semana durante la temporada de lluvias.

## 3. Reduce la velocidad

Con pavimento mojado, la distancia de frenado aumenta hasta un **40%**. La regla es simple: reduce tu velocidad al menos un 20% respecto a la que usarías en seco.

## 4. Evita charcos y baches

Los charcos ocultan baches que pueden dañar seriamente tus llantas y rines. Si no puedes evitarlos, pásalos lo más lento posible.

## 5. Revisa tus llantas después de cada lluvia fuerte

Después de manejar en lluvia intensa, revisa visualmente tus llantas buscando:

- Objetos incrustados (clavos, vidrios)
- Daños en las paredes laterales
- Pérdida de presión

## Bonus: Considera llantas para lluvia

Si vives en una zona con mucha lluvia, existen llantas diseñadas específicamente para **máxima evacuación de agua**. En ServiLlantas te asesoramos sobre la mejor opción para tu vehículo y zona.

## ¡Prepárate para las lluvias!

Visítanos en **ServiLlantas Celaya** para una revisión completa de tus llantas antes de la temporada de lluvias. ¡Tu seguridad es nuestra prioridad!`,
		Imagen:    "/images/blog/lluvias.jpg",
		Fecha:     "2024-11-20",
		Categoria: "Tips",
		Slug:      "tips-llantas-lluvias",
		Autor:     "Carlos Ramírez",
		AutorRol:  "Director General",
		Duracion:  "4 min de lectura",
		Tags:      []string{"lluvias", "seguridad", "consejos", "llantas"},
	},
	{
		ID:      4,
		Titulo:  "¿Cómo saber si mis frenos necesitan mantenimiento?",
		Resumen: "Los frenos son el sistema de seguridad más importante de tu vehículo. Aprende a identificar las señales de alerta antes de que sea tarde.",
		Contenido: `El sistema de frenos es, sin duda, el componente de seguridad más crítico de tu vehículo. Ignorar las señales de desgaste puede poner en riesgo tu vida y la de tu familia.

## Señales de alerta en tus frenos

### 1. Ruidos al frenar
- **Chirrido agudo**: Las balatas están llegando a su límite. El indicador metálico está rozando el disco.
- **Rechinido grave**: Posible daño en los discos. Necesitan revisión urgente.
- **Golpeteo**: Disco de freno deformado o dañado.

### 2. Vibración en el pedal
Si sientes pulsaciones o vibraciones al pisar el pedal de freno, es señal de que los discos están deformados y necesitan rectificación o reemplazo.

### 3. El pedal se siente esponjoso
Un pedal suave o que se va hasta el fondo indica posibles problemas con el líquido de frenos o aire en el sistema hidráulico.

### 4. El auto se jala hacia un lado
Si al frenar el vehículo se desvía, puede haber desgaste desigual en las balatas o un problema con el cáliper.

### 5. Mayor distancia de frenado
Si notas que necesitas más espacio para detenerte, tus frenos necesitan atención inmediata.

## ¿Cada cuánto revisar los frenos?

- **Balatas**: Cada 20,000-30,000 km
- **Discos**: Cada 50,000-70,000 km
- **Líquido de frenos**: Cada 2 años o 40,000 km

## No escatimes en seguridad

Los frenos no son un componente donde debas buscar el precio más bajo. En **ServiLlantas Celaya** utilizamos refacciones de calidad y ofrecemos diagnóstico gratuito de frenos. ¡Agenda tu revisión!`,
		Imagen:    "/images/blog/frenos.jpg",
		Fecha:     "2024-11-10",
		Categoria: "Seguridad",
		Slug:      "como-saber-frenos-mantenimiento",
		Autor:     "Miguel Torres",
		AutorRol:  "Jefe de Taller",
		Duracion:  "5 min de lectura",
		Tags:      []string{"frenos", "seguridad", "mantenimiento", "balatas"},
	},
	{
		ID:      5,
		Titulo:  "Guía completa del cambio de aceite: todo lo que debes saber",
		Resumen: "El cambio de aceite es el mantenimiento más básico pero más importante para tu motor. Aprende cuándo, cómo y qué aceite usar.",
		Contenido: `El aceite del motor es literalmente la sangre de tu vehículo. Lubrica, limpia, enfría y protege las piezas internas del motor. Un cambio de aceite a tiempo puede ahorrarte reparaciones costosas.

## ¿Por qué es tan importante?

El aceite cumple 4 funciones vitales:

1. **Lubricación**: Reduce la fricción entre piezas metálicas
2. **Limpieza**: Arrastra partículas y residuos de combustión
3. **Enfriamiento**: Ayuda a disipar el calor del motor
4. **Protección**: Crea una película protectora contra la corrosión

## ¿Cada cuánto cambiar el aceite?

Depende del tipo de aceite:

| Tipo de aceite | Frecuencia |
|---------------|------------|
| Mineral | Cada 5,000 km |
| Semi-sintético | Cada 7,500 km |
| Sintético | Cada 10,000-15,000 km |

## ¿Qué aceite usar?

Consulta el manual de tu vehículo para conocer la viscosidad recomendada (ejemplo: 5W-30, 10W-40). Usar el aceite incorrecto puede dañar tu motor.

### Marcas que recomendamos:
- Mobil 1
- Castrol EDGE
- Pennzoil Platinum
- Valvoline

## Señales de que necesitas un cambio

- Aceite oscuro y espeso al revisar la varilla
- Luz del aceite encendida en el tablero
- Ruidos inusuales del motor
- Humo excesivo por el escape

## En ServiLlantas te cuidamos

Cada cambio de aceite incluye:
- Aceite premium de la viscosidad correcta
- Filtro de aceite nuevo
- Revisión de 21 puntos
- Registro en tu historial de servicio

¡Visítanos en **ServiLlantas Celaya**!`,
		Imagen:    "/images/blog/aceite.jpg",
		Fecha:     "2024-10-25",
		Categoria: "Mantenimiento",
		Slug:      "guia-cambio-aceite",
		Autor:     "Carlos Ramírez",
		AutorRol:  "Director General",
		Duracion:  "6 min de lectura",
		Tags:      []string{"aceite", "motor", "mantenimiento", "lubricación"},
	},
	{
		ID:      6,
		Titulo:  "Suspensión: señales de desgaste que no debes ignorar",
		Resumen: "Una suspensión en mal estado afecta tu seguridad, comodidad y hasta el desgaste de tus llantas. Conoce las señales de alerta.",
		Contenido: `La suspensión de tu vehículo es un sistema complejo que trabaja constantemente para mantener las llantas en contacto con el pavimento, absorber las irregularidades del camino y brindarte un manejo seguro y cómodo.

## ¿Qué componentes tiene la suspensión?

- **Amortiguadores / Struts**: Controlan el rebote
- **Resortes**: Soportan el peso del vehículo
- **Rótulas**: Permiten el movimiento de la dirección
- **Terminales**: Conectan la barra de dirección con las ruedas
- **Bujes**: Amortiguan las vibraciones entre componentes
- **Barras estabilizadoras**: Reducen la inclinación en curvas

## Señales de problemas en la suspensión

### 1. Rebote excesivo
Si tu auto rebota más de 2 veces al pasar un tope, los amortiguadores están gastados.

### 2. Inclinación en curvas
Si sientes que el auto se inclina exageradamente al tomar una curva, las barras estabilizadoras o amortiguadores necesitan atención.

### 3. Ruidos al pasar baches
Golpeteos, rechinidos o tronidos al pasar por baches son señal de componentes desgastados.

### 4. Desgaste irregular de llantas
Una suspensión en mal estado causa desgaste irregular en las llantas, lo que reduce su vida útil significativamente.

### 5. Dirección imprecisa
Si sientes juego en el volante o el auto no responde con precisión, puede haber desgaste en rótulas o terminales.

## ¿Cada cuánto revisar la suspensión?

- **Amortiguadores**: Cada 60,000-80,000 km
- **Rótulas y terminales**: Revisión cada 20,000 km
- **Bujes**: Inspección visual cada 30,000 km

## Diagnóstico profesional

En **ServiLlantas Celaya** realizamos diagnóstico completo de suspensión con prueba en carretera incluida. Identificamos el problema exacto y te damos un presupuesto transparente. ¡Agenda tu cita!`,
		Imagen:    "/images/blog/suspension.jpg",
		Fecha:     "2024-10-10",
		Categoria: "Mecánica",
		Slug:      "suspension-senales-desgaste",
		Autor:     "Miguel Torres",
		AutorRol:  "Jefe de Taller",
		Duracion:  "5 min de lectura",
		Tags:      []string{"suspensión", "amortiguadores", "seguridad", "mecánica"},
	},
}

// ============ HANDLERS ============

func handleGetServicios(c *fiber.Ctx) error {
	return c.JSON(servicios)
}

func handleGetBlog(c *fiber.Ctx) error {
	// Filtrar por categoría si se envía como query param
	categoria := c.Query("categoria", "")

	if categoria != "" && categoria != "Todos" {
		var filtered []BlogPost
		for _, post := range blogPosts {
			if post.Categoria == categoria {
				filtered = append(filtered, post)
			}
		}
		return c.JSON(filtered)
	}

	return c.JSON(blogPosts)
}

func handleGetBlogPost(c *fiber.Ctx) error {
	slug := c.Params("slug")

	for _, post := range blogPosts {
		if post.Slug == slug {
			return c.JSON(post)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "Artículo no encontrado",
	})
}

func handleGetRelatedPosts(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var currentPost BlogPost
	for _, post := range blogPosts {
		if post.Slug == slug {
			currentPost = post
			break
		}
	}

	var related []BlogPost
	for _, post := range blogPosts {
		if post.Slug != slug && (post.Categoria == currentPost.Categoria) && len(related) < 3 {
			related = append(related, post)
		}
	}

	// Si no hay suficientes de la misma categoría, agregar otros
	if len(related) < 3 {
		for _, post := range blogPosts {
			if post.Slug != slug && post.Categoria != currentPost.Categoria && len(related) < 3 {
				related = append(related, post)
			}
		}
	}

	return c.JSON(related)
}

func handleContacto(c *fiber.Ctx) error {
	form := new(ContactForm)

	if err := c.BodyParser(form); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Datos inválidos.",
		})
	}

	if form.Nombre == "" || form.Telefono == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "El nombre y teléfono son obligatorios.",
		})
	}

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

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:4173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "🚗 ServiLlantas API funcionando",
		})
	})

	api.Get("/servicios", handleGetServicios)
	api.Get("/blog", handleGetBlog)
	api.Get("/blog/:slug", handleGetBlogPost)
	api.Get("/blog/:slug/related", handleGetRelatedPosts)
	api.Post("/contacto", handleContacto)

	fmt.Println("🚀 ServiLlantas API corriendo en http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
