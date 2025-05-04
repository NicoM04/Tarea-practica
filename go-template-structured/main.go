package main

import (
	"go-template/middleware"
	"go-template/routes"
	"go-template/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Inicializar conexión Mongo
	services.InitMongo()

	r := gin.Default()

	// Middleware CORS
	r.Use(middleware.CorsMiddleware())

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Leer el puerto desde el .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
