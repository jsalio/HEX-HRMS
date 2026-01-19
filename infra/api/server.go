package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hrms.local/core/contracts"
	"hrms.local/core/models"
	"hrms.local/infra/api/config"
	"hrms.local/infra/api/controller"
	"hrms.local/infra/api/middleware"
	BaseController "hrms.local/infra/api/types"
	"hrms.local/security"

	"hrms.local/repository/postgress"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router         *gin.Engine
	appController  []BaseController.Controller
	authMiddleware *middleware.AuthMiddleware
	config         *config.Config
	context        struct {
		userContract       contracts.UserContract
		departmentContract contracts.DepartmentContract
		roleContract       contracts.RoleContract
		permissionContract contracts.PermissionContract
		positionContract   contracts.PositionContract
	}
	cryptographyContext contracts.CryptographyContract
}

func NewServer(cfg *config.Config) *Server {

	server := &Server{
		router:         gin.New(),
		appController:  []BaseController.Controller{},
		authMiddleware: middleware.NewAuthMiddleware(),
		config:         cfg,
	}

	server.SetupHeaders()
	server.SetupContext()
	server.SetupControllers()

	return server
}

func (s *Server) SetupHeaders() {
	//cors config
	s.router.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		// c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
}

func (s *Server) SetupControllers() {
	s.appController = []BaseController.Controller{
		controller.NewUserController(s.authMiddleware, s.context.userContract, s.cryptographyContext),
		controller.NewRoleController(s.authMiddleware, s.context.roleContract, s.context.permissionContract),
		controller.NewDepartmentController(s.context.departmentContract, s.authMiddleware),
		controller.NewPositionController(s.context.positionContract, s.authMiddleware),
	}
}

func (s *Server) SetupContext() {
	fmt.Println("Setting up context")
	context, err := postgress.NewContext(s.config.DBURL)
	if err.Code != models.SystemErrorCodeNone {
		log.Fatal("Failed to get gorm config", err)
	}
	s.context.userContract = context.UserContract
	s.context.departmentContract = context.DepartmentContract
	s.context.roleContract = context.RoleContract
	s.context.permissionContract = context.PermissionContract
	s.context.positionContract = context.PositionContract
	s.cryptographyContext = security.NewSecurityImpl()
}

func (s *Server) StartServer() {

	for _, controller := range s.appController {
		controller.RegisterRoutes(s.router.Group("/api"))
	}

	// Register health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	srv := &http.Server{
		Addr:           ":" + s.config.ServerPort,
		Handler:        s.router,
		ReadTimeout:    s.config.ReadTimeout,
		WriteTimeout:   s.config.WriteTimeout,
		MaxHeaderBytes: s.config.MaxHeaderBytes,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
