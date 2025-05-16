package main

import (
	"github.com/Vann06/Ej_Api/db"
	"github.com/Vann06/Ej_Api/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Archivos estáticos
	r.Static("/frontend", "./frontend")

	// Conexión a la base de datos
	connection := db.ConnectDB()

	// Registrar rutas
	routes.RegisterIncidentRoutes(r, connection)

	// Levantar servidor
	r.Run(":8080")
}
