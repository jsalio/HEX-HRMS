package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	"hrms.local/core/usecases/position"
	"hrms.local/infra/api/middleware"
	"hrms.local/infra/api/types"
)

// PositionController handles HTTP requests for positions
type PositionController struct {
	*types.BaseController
	positionContract contracts.PositionContract
	authMiddleware   *middleware.AuthMiddleware
}

// NewPositionController creates a new PositionController
func NewPositionController(positionContract contracts.PositionContract, authMiddleware *middleware.AuthMiddleware) *PositionController {
	return &PositionController{
		BaseController:   types.NewBaseController("/position"),
		positionContract: positionContract,
		authMiddleware:   authMiddleware,
	}
}

// Create handles POST request to create a new position
func (pc *PositionController) Create(c *gin.Context) {
	var body models.CreatePosition
	if _, err := pc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	usecase := position.NewCreatePositionUseCase(contracts.NewGenericRequest(body), pc.positionContract)
	if err := usecase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	result, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Update handles POST request to update an existing position
func (pc *PositionController) Update(c *gin.Context) {
	var body models.ModifyPosition
	if _, err := pc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	usecase := position.NewUpdatePositionUseCase(contracts.NewGenericRequest(body), pc.positionContract)
	if err := usecase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	result, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Delete handles POST request to delete a position
func (pc *PositionController) Delete(c *gin.Context) {
	var body struct {
		ID string `json:"id"`
	}
	if _, err := pc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	filter := models.Filter{Key: "id", Value: body.ID}
	usecase := position.NewDeletePositionUseCase(pc.positionContract, contracts.NewGenericRequest(filter))
	if err := usecase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	_, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Position deleted"})
}

// Get handles POST request to get a position by ID
func (pc *PositionController) Get(c *gin.Context) {
	var body struct {
		ID string `json:"id"`
	}
	if _, err := pc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	filter := models.Filter{Key: "id", Value: body.ID}
	usecase := position.NewGetPositionByIdUseCase(pc.positionContract, contracts.NewGenericRequest(filter))
	if err := usecase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	result, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetAll handles POST request to get all positions with pagination
func (pc *PositionController) GetAll(c *gin.Context) {
	var query models.SearchQuery
	if c.Request.ContentLength > 0 {
		if _, err := pc.BaseController.GetBody(c, &query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
			return
		}
	} else {
		query = models.SearchQuery{
			Pagination: models.Pagination{Page: 1, Limit: 100},
		}
	}
	usecase := position.NewListPositionUseCase(contracts.NewGenericRequest(query), pc.positionContract)
	if err := usecase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	result, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, result)
}

// RegisterRoutes registers position routes with the router group
func (pc *PositionController) RegisterRoutes(g *gin.RouterGroup) {
	pos := g.Group("/position")
	pos.Use(pc.authMiddleware.AuthMiddleware())
	{
		pos.POST("/create", pc.Create)
		pos.POST("/update", pc.Update)
		pos.POST("/delete", pc.Delete)
		pos.POST("/get", pc.Get)
		pos.POST("/get-all", pc.GetAll)
	}
}
