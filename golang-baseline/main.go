package main

import (
	"golang-baseline/config"
	"golang-baseline/handlers"
	"golang-baseline/services"
	"golang-baseline/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize logger
	logger := utils.NewLogger()
	logger.Info("Starting Golang Baseline Web Application...")

	// Load configuration
	cfg := config.LoadConfig()
	logger.Infof("Configuration loaded - Port: %s, Environment: %s", cfg.Port, cfg.Environment)

	// Initialize service
	service := services.NewService()
	logger.Info("Service initialized")

	// Initialize handlers
	handler := handlers.NewHandler(service)
	logger.Info("Handlers initialized")

	// Setup router
	router := setupRoutes(handler)
	logger.Info("Routes configured")

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Create server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      c.Handler(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Infof("Server starting on port %s", cfg.Port)
	logger.Info("Available endpoints:")
	logger.Info("  GET  /                     - Welcome message")
	logger.Info("  GET  /health               - Health check")
	logger.Info("  GET  /api/dashboard        - Dashboard statistics")
	logger.Info("  GET  /api/backlogs         - Get all backlogs")
	logger.Info("  POST /api/backlogs         - Create backlog")
	logger.Info("  GET  /api/backlogs/{id}    - Get specific backlog")
	logger.Info("  GET  /api/stories          - Get stories by backlog")
	logger.Info("  POST /api/stories          - Create story")
	logger.Info("  GET  /api/stories/{id}     - Get specific story")
	logger.Info("  PUT  /api/stories/{id}/status - Update story status")
	logger.Info("  GET  /api/subtasks         - Get subtasks by story")
	logger.Info("  POST /api/subtasks         - Create subtask")
	logger.Info("  GET  /api/subtasks/{id}    - Get specific subtask")
	logger.Info("  PUT  /api/subtasks/{id}/status - Update subtask status")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Errorf("Could not start server: %s", err.Error())
	}
}

func setupRoutes(handler *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	// Root and health endpoints
	router.HandleFunc("/", handler.Welcome).Methods("GET")
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Dashboard
	api.HandleFunc("/dashboard", handler.GetDashboard).Methods("GET")

	// Backlog routes
	api.HandleFunc("/backlogs", handler.GetAllBacklogs).Methods("GET")
	api.HandleFunc("/backlogs", handler.CreateBacklog).Methods("POST")
	api.HandleFunc("/backlogs/{id}", handler.GetBacklog).Methods("GET")

	// Story routes
	api.HandleFunc("/stories", handler.CreateStory).Methods("POST")
	api.HandleFunc("/stories/{id}", handler.GetStory).Methods("GET")
	api.HandleFunc("/stories/{id}/status", handler.UpdateStoryStatus).Methods("PUT")
	api.HandleFunc("/backlogs/{backlogId}/stories", handler.GetStoriesByBacklog).Methods("GET")

	// SubTask routes
	api.HandleFunc("/subtasks", handler.CreateSubTask).Methods("POST")
	api.HandleFunc("/subtasks/{id}", handler.GetSubTask).Methods("GET")
	api.HandleFunc("/subtasks/{id}/status", handler.UpdateSubTaskStatus).Methods("PUT")
	api.HandleFunc("/stories/{storyId}/subtasks", handler.GetSubTasksByStory).Methods("GET")

	return router
}
