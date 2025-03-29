package main

import (
	"github.com/Vann06/Ej_Api/db"
	"github.com/Vann06/Ej_Api/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Conexión a la base de datos
	connection := db.ConnectDB()

	// Registrar las rutas y pasar la conexión
	routes.RegisterIncidentRoutes(r, connection)

	// Levantar el servidor
	r.Run(":8080")
}
