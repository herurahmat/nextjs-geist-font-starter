package handlers

import (
	"encoding/json"
	"golang-baseline/models"
	"golang-baseline/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler contains the service and handles HTTP requests
type Handler struct {
	service *services.Service
}

// NewHandler creates a new handler instance
func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Helper function to send JSON response
func (h *Handler) sendResponse(w http.ResponseWriter, statusCode int, success bool, data interface{}, errorMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: success,
		Data:    data,
		Error:   errorMsg,
	}

	json.NewEncoder(w).Encode(response)
}

// Health check endpoint
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.sendResponse(w, http.StatusOK, true, map[string]string{"status": "healthy"}, "")
}

// Backlog handlers
func (h *Handler) CreateBacklog(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBacklogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendResponse(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}

	backlog, err := h.service.CreateBacklog(req)
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusCreated, true, backlog, "")
}

func (h *Handler) GetBacklog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	backlog, err := h.service.GetBacklog(id)
	if err != nil {
		h.sendResponse(w, http.StatusNotFound, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, backlog, "")
}

func (h *Handler) GetAllBacklogs(w http.ResponseWriter, r *http.Request) {
	backlogs, err := h.service.GetAllBacklogs()
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, backlogs, "")
}

// Story handlers
func (h *Handler) CreateStory(w http.ResponseWriter, r *http.Request) {
	var req models.CreateStoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendResponse(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}

	story, err := h.service.CreateStory(req)
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusCreated, true, story, "")
}

func (h *Handler) GetStory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	story, err := h.service.GetStory(id)
	if err != nil {
		h.sendResponse(w, http.StatusNotFound, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, story, "")
}

func (h *Handler) GetStoriesByBacklog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	backlogID := vars["backlogId"]

	stories, err := h.service.GetStoriesByBacklog(backlogID)
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, stories, "")
}

// SubTask handlers
func (h *Handler) CreateSubTask(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSubTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendResponse(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}

	subtask, err := h.service.CreateSubTask(req)
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusCreated, true, subtask, "")
}

func (h *Handler) GetSubTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	subtask, err := h.service.GetSubTask(id)
	if err != nil {
		h.sendResponse(w, http.StatusNotFound, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, subtask, "")
}

func (h *Handler) GetSubTasksByStory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storyID := vars["storyId"]

	subtasks, err := h.service.GetSubTasksByStory(storyID)
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, subtasks, "")
}

// Status update handlers
func (h *Handler) UpdateStoryStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendResponse(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}

	err := h.service.UpdateStoryStatus(id, req.Status)
	if err != nil {
		h.sendResponse(w, http.StatusNotFound, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, map[string]string{"message": "Status updated successfully"}, "")
}

func (h *Handler) UpdateSubTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendResponse(w, http.StatusBadRequest, false, nil, "Invalid request body")
		return
	}

	err := h.service.UpdateSubTaskStatus(id, req.Status)
	if err != nil {
		h.sendResponse(w, http.StatusNotFound, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, map[string]string{"message": "Status updated successfully"}, "")
}

// Dashboard handler
func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.GetDashboardStats()
	if err != nil {
		h.sendResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	h.sendResponse(w, http.StatusOK, true, stats, "")
}

// Welcome handler
func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	welcomeData := map[string]interface{}{
		"message":     "Welcome to Golang Baseline Web Application",
		"description": "Project Management and Backlog Tracking System",
		"version":     "1.0.0",
		"endpoints": map[string]string{
			"health":     "/health",
			"dashboard":  "/api/dashboard",
			"backlogs":   "/api/backlogs",
			"stories":    "/api/stories",
			"subtasks":   "/api/subtasks",
		},
	}

	h.sendResponse(w, http.StatusOK, true, welcomeData, "")
}
