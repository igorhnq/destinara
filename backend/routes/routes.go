package routes

import (
	"destinara-backend/controllers"
	"destinara-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Middleware CORS
	router.Use(middleware.CORSMiddleware())

	// Rotas públicas
	public := router.Group("/api")
	{
		// Autenticação
		auth := public.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Pacotes de viagem (públicos)
		packages := public.Group("/packages")
		{
			packages.GET("/", controllers.GetAllPackages)
			packages.GET("/:id", controllers.GetPackageByID)
		}
	}

	// Rotas protegidas
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Pacotes de viagem (admin)
		adminPackages := protected.Group("/admin/packages")
		{
			adminPackages.POST("/", controllers.CreatePackage)
			adminPackages.PUT("/:id", controllers.UpdatePackage)
			adminPackages.DELETE("/:id", controllers.DeletePackage)
		}

		// Reservas
		bookings := protected.Group("/bookings")
		{
			bookings.GET("/", controllers.GetUserBookings)
			bookings.GET("/:id", controllers.GetBookingByID)
			bookings.POST("/", controllers.CreateBooking)
			bookings.PUT("/:id", controllers.UpdateBooking)
			bookings.DELETE("/:id", controllers.CancelBooking)
		}
	}

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Destinara API está funcionando",
		})
	})

	return router
} 