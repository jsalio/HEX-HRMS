package types

import (
	"net/http"

	"hrms.local/core/models"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.RouterGroup)
}

type BaseController struct {
	Path string
}

func NewBaseController(path string) *BaseController {
	return &BaseController{
		Path: path,
	}
}

func (bc *BaseController) GetBody(c *gin.Context, target interface{}) (interface{}, *models.SystemError) {
	if err := c.ShouldBindJSON(&target); err != nil {
		return nil, &models.SystemError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		}
	}
	return target, nil
}
