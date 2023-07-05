package handler

import (
	"FinalBack/internal/paciente"
	"FinalBack/internal/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
    PacienteService paciente.PacienteService
}

func NewPacienteHandler(pacienteService paciente.PacienteService) *PacienteHandler {
    return &PacienteHandler{
        PacienteService: pacienteService,
    }
}

func (h *PacienteHandler) RegisterRoutes(router *gin.RouterGroup) {
    router.GET("/paciente/:id", h.GetOne)
    router.POST("/paciente", h.Create)
    router.PUT("/paciente/:id", h.UpdateOne)
    router.DELETE("/paciente/:id", h.DeleteOne)
	router.PATCH("/paciente/:id", h.UpdateField)
}

//localhost:8080/api/paciente/3
func (h *PacienteHandler) GetOne(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
        return
    }

    paciente, err := h.PacienteService.GetPacienteById(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "paciente not found"})
        return
    }

    ctx.JSON(http.StatusOK, paciente)
}

func (h *PacienteHandler) Create(ctx *gin.Context) {
    var paciente domain.Paciente
    err := ctx.BindJSON(&paciente)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
        return
    }

    result := h.PacienteService.CreatePaciente(&paciente)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not create paciente"})
        return
    }

    ctx.JSON(http.StatusCreated, result)
}

func (h *PacienteHandler) UpdateOne(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
        return
    }

    var paciente domain.Paciente
    err = ctx.BindJSON(&paciente)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
        return
    }
    paciente.Id = id

    err = h.PacienteService.UpdatePaciente(&paciente)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not update paciente"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "paciente updated successfully"})
}

func (h *PacienteHandler) DeleteOne(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.PacienteService.DeletePaciente(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete paciente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "paciente deleted"})
}

func (h *PacienteHandler) UpdateField(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
        return
    }

    var updateData struct {
        Field string      `json:"field" binding:"required"`
        Value interface{} `json:"value" binding:"required"`
    }
    err = c.ShouldBindJSON(&updateData)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
        return
    }

    err = h.PacienteService.UpdatePacienteField(id, updateData.Field, updateData.Value)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update patient: %v", err)})
        return
    }

    c.Status(http.StatusOK)
}