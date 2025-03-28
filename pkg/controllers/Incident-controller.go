package controllers

import (
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/Vann06/Ej_Api/pkg/models"
	"github.com/gin-gonic/gin"
)

var incidents = []models.Incident{}

// POST - crear incidentes
func CreateIncident(c *gin.Context) {
	var newIncident models.Incident

	if err := c.BindJSON(&newIncident); err != nil {
		return
	}

	if len(newIncident.Description) < 10 || newIncident.Reporter == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos incompletos"})
		return
	}

	newIncident.ID = len(incidents) + 1
	newIncident.Status = "pendiente"
	newIncident.CreatedAt = time.Now()

	incidents = append(incidents, newIncident)

	c.IndentedJSON(http.StatusCreated, newIncident)
}

// GET - lista de incidentes
func GetIncidents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, incidents)
}

// GET - incidente por ID
func GetIncidentByID(c *gin.Context) {
	id := c.Param("id")

	// loop para buscar por ID
	for _, incident := range incidents {
		if strconv.Itoa(incident.ID) == id{
			c.JSON(http.StatusOK, incident)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})

}

// PUT incidentes - actualizar el estado 
func UpdateIncidents(c *gin.Context){

	id := c.Param("id")
	
	// recibir el campo modificable de status
	var input struct {
		Status string `json:"status"`
	}

	statusValidos := map[string]bool{
		"pendiente": true,
		"en proceso": true,
		"resuelto": true,
	}

	// Si el input no cuenta como estatus 
	if !statusValidos[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estado no valido"})
		return
	}

	for i, incident := range incidents {
		if strconv.Itoa(incident.ID) == id{
			incidents[i].Status = input.Status
			c.JSON(http.StatusOK, gin.H{"message": "Estado actualizado"})
			return
		}
	}
	// si no lo encontro 
	c.JSON(http.StatusNotFound,  gin.H{"error": "Incidente no encontrado"})

}

// DELETE eliminar incidentes 
func DeleteIncidents(c *gin.Context){
	
	id := c.Param("id")

	for i, incident := range incidents {
		if strconv.Itoa(incident.ID) == id{
			// corta y junta los incidentes eliminando el que se desea
			incidents = slices.Delete(incidents, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "Incidente eliminado"})
			return
		}
	}
	c.JSON(http.StatusNotFound,  gin.H{"error": "Incidente no encontrado"})
}