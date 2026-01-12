package controller

import (
	"net/http"

	"hrms.local/core/contracts"
	"hrms.local/core/models"
	permissionUseCase "hrms.local/core/usecases/permissions"
	roleUseCase "hrms.local/core/usecases/roles"
	"hrms.local/infra/api/middleware"
	"hrms.local/infra/api/types"
	"hrms.local/repository/postgress/repo"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	*types.BaseController
	roleContract       contracts.RoleContract
	permissionContract contracts.PermissionContract
	authMiddleware     *middleware.AuthMiddleware
}

func NewRoleController(authMiddleware *middleware.AuthMiddleware, roleContract contracts.RoleContract, permissionContract contracts.PermissionContract) *RoleController {
	return &RoleController{
		BaseController:     types.NewBaseController("/roles"),
		roleContract:       roleContract,
		permissionContract: permissionContract,
		authMiddleware:     authMiddleware,
	}
}

func (rc *RoleController) SetContext(c *gin.Context) {
	if r, ok := rc.roleContract.(*repo.RoleRepository); ok {
		r.WithContext(c.Request.Context())
	}
}

func (rc *RoleController) Create(c *gin.Context) {
	rc.SetContext(c)
	var body models.CreateRole
	if _, err := rc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	useCase := roleUseCase.NewCreateRoleUsecase(rc.roleContract, contracts.NewGenericRequest(body))
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	createdRole, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, createdRole)
}

func (rc *RoleController) Update(c *gin.Context) {
	rc.SetContext(c)
	var body models.Role
	if _, err := rc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	useCase := roleUseCase.NewUpdateRoleUsecase(rc.roleContract, contracts.NewGenericRequest(body))
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	updatedRole, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, updatedRole)
}

func (rc *RoleController) Delete(c *gin.Context) {
	rc.SetContext(c)
	var body struct {
		ID string `json:"id"`
	}
	if _, err := rc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	useCase := roleUseCase.NewDeleteRoleUsecase(rc.roleContract, body.ID)
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	if err := useCase.Execute(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted"})
}

func (rc *RoleController) Get(c *gin.Context) {
	rc.SetContext(c)
	var body struct {
		ID string `json:"id"`
	}
	if _, err := rc.BaseController.GetBody(c, &body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	useCase := roleUseCase.NewGetRoleUsecase(rc.roleContract, body.ID)
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	role, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (rc *RoleController) GetAll(c *gin.Context) {
	rc.SetContext(c)
	var query models.SearchQuery
	if c.Request.ContentLength > 0 {
		if _, err := rc.BaseController.GetBody(c, &query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
			return
		}
	} else {
		query = models.SearchQuery{
			Pagination: models.Pagination{Page: 1, Limit: 100},
		}
	}

	useCase := roleUseCase.NewListRolesUsecase(rc.roleContract, contracts.NewGenericRequest(query))
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	roles, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (rc *RoleController) GetPermissions(c *gin.Context) {
	rc.SetContext(c)
	roleID := c.Param("role_id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role_id is required"})
		return
	}

	useCase := roleUseCase.NewGetPermissionsUsecase(rc.roleContract, roleID)
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		return
	}

	permissions, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, permissions)
}

func (rc *RoleController) ListSystemPermissions(c *gin.Context) {
	useCase := permissionUseCase.NewListPermissionsUseCase(rc.permissionContract)
	permissions, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, permissions)
}

func (rc *RoleController) RegisterRoutes(router *gin.RouterGroup) {
	roles := router.Group("/roles")
	roles.Use(rc.authMiddleware.AuthMiddleware())
	{
		roles.POST("/create", rc.Create)
		roles.POST("/update", rc.Update)
		roles.POST("/delete", rc.Delete)
		roles.POST("/get", rc.Get)
		roles.POST("/get-all", rc.GetAll)
		roles.GET("/get-permissions/:role_id", rc.GetPermissions)
		roles.GET("/system-permissions", rc.ListSystemPermissions)
	}
}
