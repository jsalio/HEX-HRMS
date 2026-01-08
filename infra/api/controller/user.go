package controller

import (
	"hrms/core/contracts"
	"hrms/core/models"
	userUseCase "hrms/core/usecases/users"
	"hrms/infra/api/middleware"
	"hrms/infra/api/types"
	"hrms/repository/postgress/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*types.BaseController
	userContract   contracts.UserContract
	authMiddleware *middleware.AuthMiddleware
}

func NewUserController(authMiddleware *middleware.AuthMiddleware, userContract contracts.UserContract) *UserController {
	return &UserController{
		BaseController: types.NewBaseController("/auth"),
		userContract:   userContract,
		authMiddleware: authMiddleware,
	}
}

func (uc *UserController) SetContext(c *gin.Context) {
	if r, ok := uc.userContract.(*repo.UserRepository); ok {
		r.WithContext(c.Request.Context())
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var body models.CreateUser
	uc.SetContext(c)
	_, err := uc.BaseController.GetBody(c, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		c.Abort()
		return
	}

	useCase := userUseCase.NewCreateUserUseCase(uc.userContract, contracts.NewGenericRequest(body))
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	data, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var body models.LoginUser
	uc.SetContext(c)
	_, err := uc.BaseController.GetBody(c, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	useCase := userUseCase.NewLoginUserUseCase(uc.userContract, contracts.NewGenericRequest(body))
	if err := useCase.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	data, err := useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	tokenData, err2 := uc.authMiddleware.GenerateToken(data.Username, map[string]interface{}{
		"username": data.Username,
		"type":     data.Type,
		"email":    data.Email,
	})
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenData})
}

func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	uc.authMiddleware.Config.AddPublicRoute("POST", "/api/auth")
	public := router.Group("/auth")
	{
		public.POST("/", uc.CreateUser)
		public.POST("/login", uc.LoginUser)
	}
	//return router
}
