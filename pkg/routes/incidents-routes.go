package routes

import (
	"database/sql"

	"github.com/Vann06/Ej_Api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterIncidentRoutes(r *gin.Engine, db *sql.DB) {
	incidentController := controllers.NewIncidentController(db)

	incidents := r.Group("/incidents")
	{
		incidents.POST("", incidentController.CreateIncident)
		incidents.GET("", incidentController.GetIncidents)
		incidents.GET("/:id", incidentController.GetIncidentByID)
		incidents.PUT("/:id", incidentController.UpdateIncidentStatus)
		incidents.DELETE("/:id", incidentController.DeleteIncident)
	}
}
