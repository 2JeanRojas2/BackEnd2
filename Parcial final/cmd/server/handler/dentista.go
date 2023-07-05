package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"parcialfinal/internal/domain"
	"parcialfinal/internal/domain/dentista"
	"parcialfinal/pkg/web"
)

type DentistaHandler struct {
	DentistaService dentista.IService
}

func NewDentistaHandler(dentistaService dentista.IService) *DentistaHandler {
	return &DentistaHandler{
		DentistaService: dentistaService,
	}
}

func (h *DentistaHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/dentistas/:id", h.GetOne)
	router.POST("/dentistas", h.Create)
	router.PUT("/dentistas/:id", h.UpdateOne)
	router.DELETE("/dentistas/:id", h.DeleteOne)
}

func (h *DentistaHandler) GetOne(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id parameter")
		return
	}

	dentista, err := h.DentistaService.GetDentistaBy(id)
	if err != nil {
		web.NewBadRequestApiError("dentista_id not found")
		return
	}

	ctx.JSON(http.StatusOK, dentista)
}

func (h *DentistaHandler) Create(ctx *gin.Context) {
	var dentista domain.Dentista
	err := ctx.BindJSON(&dentista)
	if err != nil {
		web.NewBadRequestApiError("invalid JSON body")
		return
	}

	result, err := h.DentistaService.CreateDentista(dentista)
	if err != nil {
		web.NewBadRequestApiError("It was not possible to create the dentist")
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (h *DentistaHandler) UpdateOne(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id parameter")
		return
	}

	var dentista domain.Dentista
	err = ctx.BindJSON(&dentista)
	if err != nil {
		web.NewBadRequestApiError("invalid JSON body")
		return
	}
	dentista.Id = id

	err = h.DentistaService.UpdateDentista(dentista)
	if err != nil {
		web.NewBadRequestApiError("Could not update dentist")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dentista updated successfully"})
}

func (h *DentistaHandler) DeleteOne(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id parameter")
		return
	}

	err = h.DentistaService.DeleteDentistaBy(id)
	if err != nil {
		web.NewBadRequestApiError("Could not delete dentist")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dentista deleted successfully"})
}