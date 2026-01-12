package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hrms.local/core/contracts"
	"hrms.local/core/models"
	"hrms.local/core/usecases/department"
	"hrms.local/infra/api/middleware"
	"hrms.local/infra/api/types"
)

type DepartmentController struct {
	*types.BaseController
	departmentContract contracts.DepartmentContract
	authMiddleware     *middleware.AuthMiddleware
}

func NewDepartmentController(departmentContract contracts.DepartmentContract, authMiddleware *middleware.AuthMiddleware) *DepartmentController {
	return &DepartmentController{
		BaseController:     types.NewBaseController("/department"),
		departmentContract: departmentContract,
		authMiddleware:     authMiddleware,
	}
}

func (dc *DepartmentController) Create(c *gin.Context) {
	var body models.Department
	if _, err := dc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	usecase := department.NewCreateDepartmentUseCase(contracts.NewGenericRequest(body), dc.departmentContract)
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

func (dc *DepartmentController) Update(c *gin.Context) {
	var body models.Department
	if _, err := dc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	usecase := department.NewUpdateDepartmentUseCase(contracts.NewGenericRequest(body), dc.departmentContract)
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

func (dc *DepartmentController) Delete(c *gin.Context) {
	var body struct {
		ID string `json:"id"`
	}
	if _, err := dc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	filter := models.Filter{Key: "id", Value: body.ID}
	usecase := department.NewDeleteDepartmentUsecase(dc.departmentContract, contracts.NewGenericRequest(filter))
	if err := usecase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	_, err := usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted"})
}

func (dc *DepartmentController) Get(c *gin.Context) {
	var body struct {
		ID string `json:"id"`
	}
	if _, err := dc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}
	filter := models.Filter{Key: "id", Value: body.ID}
	usecase := department.NewGetDepartmentByIdUseCase(dc.departmentContract, contracts.NewGenericRequest(filter))
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

func (dc *DepartmentController) GetAll(c *gin.Context) {
	var query models.SearchQuery
	if c.Request.ContentLength > 0 {
		if _, err := dc.BaseController.GetBody(c, &query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
			return
		}
	} else {
		query = models.SearchQuery{
			Pagination: models.Pagination{Page: 1, Limit: 100},
		}
	}
	usecase := department.NewListDepartmentUseCase(contracts.NewGenericRequest(query), dc.departmentContract)
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

func (dc *DepartmentController) RegisterRoutes(g *gin.RouterGroup) {
	dept := g.Group("/department")
	dept.Use(dc.authMiddleware.AuthMiddleware())
	{
		dept.POST("/create", dc.Create)
		dept.POST("/update", dc.Update)
		dept.POST("/delete", dc.Delete)
		dept.POST("/get", dc.Get)
		dept.POST("/get-all", dc.GetAll)
	}
}
