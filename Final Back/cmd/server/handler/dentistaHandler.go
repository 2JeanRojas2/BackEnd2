package handler

import (
	"FinalBack/internal/dentista"
	"FinalBack/internal/domain"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistaHandler struct {
    DentistaService dentista.DentistService
}

func NewDentistaHandler(dentistaService dentista.DentistService) *DentistaHandler {
    return &DentistaHandler{
        DentistaService: dentistaService,
    }
}

func (h *DentistaHandler) RegisterRoutes(router *gin.RouterGroup) {
    router.GET("/dentistas/:id", h.GetOne)
    router.POST("/dentistas", h.Create)
    router.PUT("/dentistas/:id", h.UpdateOne)
    router.DELETE("/dentistas/:id", h.DeleteOne)
	router.PATCH("/dentistas/:id", h.UpdateField)
}

//localhost:8080/api/dentistas/1
func (h *DentistaHandler) GetOne(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
        return
    }

    dentista, err := h.DentistaService.GetDentistaById(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "dentista not found"})
        return
    }

    ctx.JSON(http.StatusOK, dentista)
}

// localhost:8080/api/dentistas
func (h *DentistaHandler) Create(ctx *gin.Context) {
    var dentista domain.Dentista
    err := ctx.BindJSON(&dentista)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
        return
    }

    result := h.DentistaService.CreateDentist(&dentista)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not create dentista"})
        return
    }

    ctx.JSON(http.StatusCreated, result)
}

//localhost:8080/api/dentistas/7
func (h *DentistaHandler) UpdateOne(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
        return
    }

    var dentista domain.Dentista
    err = ctx.BindJSON(&dentista)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
        return
    }
    dentista.Id = id

    err = h.DentistaService.UpdateDentist(&dentista)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not update dentista"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "dentista updated successfully"})
}

func (h *DentistaHandler) UpdateField(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return
    }

    field := c.Param("field")

    var value interface{}
    switch field {
    case "nombre", "apellido":
        value = c.PostForm("value")
    case "matricula":
        value = c.PostForm("value")
        if _, err := strconv.Atoi(value.(string)); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for matricula"})
            return
        }
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field"})
        return
    }

    err = h.DentistaService.UpdateDentistField(id, field, value)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Dentist updated successfully"})
}

func (h *DentistaHandler) DeleteOne(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
        return
    }

    err = h.DentistaService.DeleteDentist(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not delete dentista"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "dentista deleted successfully"})
}


