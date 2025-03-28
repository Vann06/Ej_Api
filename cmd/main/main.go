package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Vann06/Ej_Api/pkg/routes"
)

func main() {
	r := gin.Default()

	// Registrar las rutas
	routes.RegisterIncidentRoutes(r)

	// usar el servidor 8080
	r.Run(":8080")
}
