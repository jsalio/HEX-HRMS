package main

import (
	"context"
	"hrms/core/contracts"
	"hrms/infra/api/config"
	"hrms/infra/api/controller"
	"hrms/infra/api/middleware"
	BaseController "hrms/infra/api/types"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hrms/repository/postgress/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	server := NewServer(cfg)
	server.StartServer()
}

type Server struct {
	router         *gin.Engine
	appController  []BaseController.Controller
	authMiddleware *middleware.AuthMiddleware
	config         *config.Config
	context        struct {
		userContract contracts.UserContract
	}
}

func NewServer(cfg *config.Config) *Server {

	server := &Server{
		router:         gin.New(),
		appController:  []BaseController.Controller{},
		authMiddleware: middleware.NewAuthMiddleware(),
		config:         cfg,
	}

	server.SetupHeaders()
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
		controller.NewUserController(s.authMiddleware, s.context.userContract),
	}
}

func (s *Server) SetupContext(db any) {
	userContract := repo.NewUserRepository(nil)
	// userContract.(*repo.GenericCrud[models.User]).WithContext(context.Background()) //use this in case you need to set the context
	s.context.userContract = userContract
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
