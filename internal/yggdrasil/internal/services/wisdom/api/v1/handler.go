package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/domain"
	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/services/wisdom/api/v1/dto"
	"github.com/ognick/job-interview-playground/pkg/logger"
)

type usecase interface {
	GetWisdom(ctx context.Context) (domain.Wisdom, error)
}

type Handler struct {
	log     logger.Logger
	usecase usecase
}

func NewHandler(
	log logger.Logger,
	usecase usecase,
) *Handler {
	return &Handler{
		log:     log,
		usecase: usecase,
	}
}

func (h *Handler) getWisdom(c *gin.Context) {
	wisdom, err := h.usecase.GetWisdom(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, dto.NewWisdom(wisdom))
}

func (h *Handler) Register(router gin.IRouter) {
	v1 := router.Group("/v1")
	{
		v1.GET("/wisdom", h.getWisdom)
	}
}
