package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Vann06/Ej_Api/pkg/models"
	"github.com/gin-gonic/gin"
)

type IncidentController struct {
	Db *sql.DB
}

func NewIncidentController(db *sql.DB) *IncidentController {
	return &IncidentController{Db: db}
}

// esto es para el POST
func (ic *IncidentController) CreateIncident(c *gin.Context) {
	var newIncident models.Incident

	if err := c.ShouldBindJSON(&newIncident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Validaciones de negocio
	if len(newIncident.Description) < 10 || newIncident.Reporter == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos incompletos o inválidos"})
		return
	}

	// Insertar en la base de datos
	result, err := ic.Db.Exec(`
		INSERT INTO ticket (reporter, description, status)
		VALUES (?, ?, ?)`,
		newIncident.Reporter, newIncident.Description, newIncident.Status,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar en la base de datos"})
		return
	}

	insertedID, _ := result.LastInsertId()
	newIncident.ID = int(insertedID)
	c.JSON(http.StatusCreated, newIncident)
}

// esto es para el GET de todos los incidentes
func (ic *IncidentController) GetIncidents(c *gin.Context) {
	rows, err := ic.Db.Query("SELECT id, reporter, description, status, created_at FROM ticket")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}
	defer rows.Close()

	var incidents []models.Incident

	for rows.Next() {
		var inc models.Incident
		err := rows.Scan(&inc.ID, &inc.Reporter, &inc.Description, &inc.Status, &inc.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer datos"})
			return
		}
		incidents = append(incidents, inc)
	}

	c.JSON(http.StatusOK, incidents)
}

// esto es para el GET con ID
func (ic *IncidentController) GetIncidentByID(c *gin.Context) {
	idParam := c.Param("id")

	// Convertir string a int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var incident models.Incident

	// Consulta a la base de datos
	err = ic.Db.QueryRow(`
		SELECT id, reporter, description, status, created_at
		FROM ticket
		WHERE id = ?
	`, id).Scan(&incident.ID, &incident.Reporter, &incident.Description, &incident.Status, &incident.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Incidente no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		}
		return
	}

	c.JSON(http.StatusOK, incident)
}

// esto es para PUT

func (ic *IncidentController) UpdateIncidentStatus(c *gin.Context) {
	idParam := c.Param("id")

	// Convertir el ID a int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Estructura para recibir el campo "status"
	var input struct {
		Status string `json:"status"`
	}

	// Validar el body JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Validar si el status es uno permitido
	validStatus := map[string]bool{
		"pendiente":  true,
		"en proceso": true,
		"resuelto":   true,
	}

	if !validStatus[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estado no válido"})
		return
	}

	// Ejecutar UPDATE
	result, err := ic.Db.Exec(`
		UPDATE ticket
		SET status = ?
		WHERE id = ?
	`, input.Status, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar incidente"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incidente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado actualizado correctamente"})
}

// esto es para DELETE
func (ic *IncidentController) DeleteIncident(c *gin.Context) {
	idParam := c.Param("id")

	// Convertir ID a int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Ejecutar DELETE en la base de datos
	result, err := ic.Db.Exec(`
		DELETE FROM ticket
		WHERE id = ?
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar incidente"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incidente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Incidente eliminado correctamente"})
}
