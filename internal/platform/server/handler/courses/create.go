package courses

import (
	"errors"
	"net/http"

	mooc "api/internal"
	"api/internal/creating"
	"api/kit/command"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Crea la request con nuestro DTO
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// Ejecuta el caso de uso de creacion de cursos
		err := commandBus.Dispatch(ctx, creating.NewCourseCommand(
			req.ID, req.Name, req.Duration,
		))

		// Devuelve y hace handling de possibles tipos de errores
		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrEmptyCourseName), errors.Is(err, mooc.ErrEmptyDuration):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
