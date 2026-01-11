package controller

import (
	"net/http"

	"hrms.local/core/contracts"
	"hrms.local/core/models"
	userUseCase "hrms.local/core/usecases/users"
	"hrms.local/infra/api/middleware"
	"hrms.local/infra/api/types"
	"hrms.local/repository/postgress/repo"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*types.BaseController
	userContract         contracts.UserContract
	authMiddleware       *middleware.AuthMiddleware
	cryptographyContract contracts.CryptographyContract
}

func NewUserController(authMiddleware *middleware.AuthMiddleware, userContract contracts.UserContract, cryptographyContract contracts.CryptographyContract) *UserController {
	return &UserController{
		BaseController:       types.NewBaseController("/auth"),
		userContract:         userContract,
		authMiddleware:       authMiddleware,
		cryptographyContract: cryptographyContract,
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

	useCase := userUseCase.NewCreateUserUseCase(uc.userContract, contracts.NewGenericRequest(body), uc.cryptographyContract)
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

	useCase := userUseCase.NewLoginUserUseCase(uc.userContract, contracts.NewGenericRequest(body), uc.cryptographyContract)
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
	response := gin.H{
		"username": data.Username,
		"type":     data.Type,
		"email":    data.Email,
		"picture":  "",
		"role":     "",
		"token":    tokenData,
	}
	c.JSON(http.StatusOK, response)
}

func (uc *UserController) LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (uc *UserController) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Me successful"})
}

func (uc *UserController) ListUsers(c *gin.Context) {
	uc.SetContext(c)
	var body models.SearchQuery

	if c.Request.ContentLength > 0 {
		if _, err := uc.BaseController.GetBody(c, &body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Message})
			c.Abort()
			return
		}
	} else {
		body = models.SearchQuery{
			Filters: models.Filters{},
			Pagination: models.Pagination{
				Page:  1,
				Limit: 10,
			},
		}
	}

	request := contracts.NewGenericRequest(body)
	users := userUseCase.NewListUserUseCase(uc.userContract, request)
	if err := users.Validate(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	data, err := users.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
}

func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	uc.authMiddleware.Config.AddPublicRoute("POST", "/api/auth")
	uc.authMiddleware.Config.AddPublicRoute("POST", "/api/auth/login")
	routeController := router.Group("/auth")
	public := routeController.Group("/")
	{
		public.POST("/login", uc.LoginUser)
		public.POST("/", uc.CreateUser)
	}
	private := router.Group("/auth")
	private.Use(uc.authMiddleware.AuthMiddleware())
	{
		private.GET("/me", uc.Me)
		private.POST("/list", uc.ListUsers)

	}
	//return router
}
