package handler

import (
	"FinalBack/internal/turno"
	"github.com/gin-gonic/gin"
	"FinalBack/internal/domain"
	"net/http"
	"strconv"
)

type turnoHandler struct {
	turnoService turno.TurnoService
}

func NewTurnoHandler(turnoService turno.TurnoService) *turnoHandler {
	return &turnoHandler{turnoService}
}

func (h *turnoHandler) RegisterRoutes(router *gin.RouterGroup) {
    router.GET("/turno/:id", h.GetById)
	router.GET("/turnoDNI/:dni", h.GetTurnoByDniPaciente)
    router.POST("/turno", h.Create)
    router.PUT("/turno/:id", h.Update)
    router.DELETE("/turno/:id", h.Delete)
	router.PATCH("/turno/:id", h.UpdateField)
}

func (h *turnoHandler) Create(c *gin.Context) {
	var turno domain.TurnoPost
	if err := c.ShouldBindJSON(&turno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.turnoService.Create(&turno); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": turno.Id})
}

func (h *turnoHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	turno, err := h.turnoService.GetById(id)
	if err != nil {
		if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Turno not found"})
        return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, turno)
}

func (h *turnoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var turno domain.TurnoPost
	if err := c.ShouldBindJSON(&turno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	turno.Id = id

	if err := h.turnoService.Update(&turno); err != nil {
		if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Turno not found"})
        return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *turnoHandler) UpdateField(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	field := c.Param("field")
	value := c.Param("value")

	if err := h.turnoService.UpdateTurnoField(id, field, value); err != nil {
		if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "turno not found"})
        return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *turnoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.turnoService.Delete(id)
	if err != nil {
		if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "turno not found"})
        return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting turno"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Turno deleted"})
}

func (h *turnoHandler) GetTurnoByDniPaciente(c *gin.Context) {
	dni := c.Param("dni")

turnos, err := h.turnoService.GetTurnoByDniPaciente(dni)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron turnos para el paciente con el DNI especificado"})
		return
	}

	var response []map[string]interface{}
	for _, turno := range turnos {
		response = append(response, map[string]interface{}{
			"id":           turno.Id,
			"FechaYHora":   turno.FechaYHora,
			"Descripcion":  turno.Descripcion,
			"Paciente": map[string]interface{}{
				"id":        turno.Paciente.Id,
				"nombre":    turno.Paciente.Nombre,
				"apellido":  turno.Paciente.Apellido,
				"domicilio": turno.Paciente.Domicilio,
				"dni":       turno.Paciente.Dni,
			},
			"Dentista": map[string]interface{}{
				"id":        turno.Dentista.Id,
				"nombre":    turno.Dentista.Nombre,
				"apellido":  turno.Dentista.Apellido,
				"matricula": turno.Dentista.Matricula,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"turnos": response})
}
