package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Vann06/Ej_Api/pkg/controllers"

)

func RegisterIncidentRoutes(r *gin.Engine) {
	incidents := r.Group("/incidents")
	{
		incidents.GET("", controllers.GetIncidents)
		incidents.GET("/:id", controllers.GetIncidentByID)
		incidents.POST("", controllers.CreateIncident)
		incidents.PUT("/:id", controllers.UpdateIncidents)
		incidents.DELETE("/:id", controllers.DeleteIncidents)
	}
}
