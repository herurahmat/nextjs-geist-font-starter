package models

import (
	"time"
)

// Status represents the status of a backlog, story, or task
type Status string

const (
	StatusTodo       Status = "TODO"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
	StatusBlocked    Status = "BLOCKED"
)

// Backlog represents a project backlog
type Backlog struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Stories     []Story   `json:"stories"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Story represents a user story within a backlog
type Story struct {
	ID           string    `json:"id"`
	BacklogID    string    `json:"backlog_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	JiraURL      string    `json:"jira_url"`
	EffortOrigin int       `json:"effort_origin"`
	PIC          string    `json:"pic"`
	PlanStart    time.Time `json:"plan_start"`
	PlanEnd      time.Time `json:"plan_end"`
	ActualStart  *time.Time `json:"actual_start,omitempty"`
	ActualEnd    *time.Time `json:"actual_end,omitempty"`
	Status       Status    `json:"status"`
	SubTasks     []SubTask `json:"subtasks"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// SubTask represents a subtask within a story
type SubTask struct {
	ID          string     `json:"id"`
	StoryID     string     `json:"story_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Effort      int        `json:"effort"`
	JiraURL     string     `json:"jira_url"`
	PIC         string     `json:"pic"`
	PlanStart   time.Time  `json:"plan_start"`
	PlanEnd     time.Time  `json:"plan_end"`
	ActualStart *time.Time `json:"actual_start,omitempty"`
	ActualEnd   *time.Time `json:"actual_end,omitempty"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// CreateBacklogRequest represents the request to create a new backlog
type CreateBacklogRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

// CreateStoryRequest represents the request to create a new story
type CreateStoryRequest struct {
	BacklogID    string    `json:"backlog_id" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	Description  string    `json:"description"`
	JiraURL      string    `json:"jira_url"`
	EffortOrigin int       `json:"effort_origin"`
	PIC          string    `json:"pic"`
	PlanStart    time.Time `json:"plan_start"`
	PlanEnd      time.Time `json:"plan_end"`
}

// CreateSubTaskRequest represents the request to create a new subtask
type CreateSubTaskRequest struct {
	StoryID     string    `json:"story_id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Effort      int       `json:"effort"`
	JiraURL     string    `json:"jira_url"`
	PIC         string    `json:"pic"`
	PlanStart   time.Time `json:"plan_start"`
	PlanEnd     time.Time `json:"plan_end"`
}

// UpdateStatusRequest represents the request to update status
type UpdateStatusRequest struct {
	Status Status `json:"status" validate:"required"`
}
