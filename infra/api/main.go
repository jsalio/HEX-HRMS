package main

import (
	"context"
	"hrms/infa/api/controller"
	"hrms/infa/api/middleware"
	BaseController "hrms/infa/api/types"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := NewServer()
	server.StartServer()
}

type Server struct {
	router         *gin.Engine
	appController  []BaseController.Controller
	authMiddleware *middleware.AuthMiddleware
}

func NewServer() *Server {

	server := &Server{
		router:         gin.New(),
		appController:  []BaseController.Controller{},
		authMiddleware: middleware.NewAuthMiddleware(),
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
		controller.NewUserController(s.authMiddleware, nil),
	}
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
		Addr:           ":5000",
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
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
