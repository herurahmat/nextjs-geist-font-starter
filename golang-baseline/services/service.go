package services

import (
	"errors"
	"golang-baseline/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for the application
type Service struct {
	backlogs map[string]*models.Backlog
	stories  map[string]*models.Story
	subtasks map[string]*models.SubTask
	mutex    sync.RWMutex
}

// NewService creates a new service instance
func NewService() *Service {
	return &Service{
		backlogs: make(map[string]*models.Backlog),
		stories:  make(map[string]*models.Story),
		subtasks: make(map[string]*models.SubTask),
	}
}

// Backlog operations
func (s *Service) CreateBacklog(req models.CreateBacklogRequest) (*models.Backlog, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	backlog := &models.Backlog{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Stories:     []models.Story{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.backlogs[backlog.ID] = backlog
	return backlog, nil
}

func (s *Service) GetBacklog(id string) (*models.Backlog, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	backlog, exists := s.backlogs[id]
	if !exists {
		return nil, errors.New("backlog not found")
	}

	// Load stories for this backlog
	var stories []models.Story
	for _, story := range s.stories {
		if story.BacklogID == id {
			// Load subtasks for this story
			var subtasks []models.SubTask
			for _, subtask := range s.subtasks {
				if subtask.StoryID == story.ID {
					subtasks = append(subtasks, *subtask)
				}
			}
			story.SubTasks = subtasks
			stories = append(stories, *story)
		}
	}
	backlog.Stories = stories

	return backlog, nil
}

func (s *Service) GetAllBacklogs() ([]*models.Backlog, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var backlogs []*models.Backlog
	for _, backlog := range s.backlogs {
		// Load stories for each backlog
		var stories []models.Story
		for _, story := range s.stories {
			if story.BacklogID == backlog.ID {
				// Load subtasks for this story
				var subtasks []models.SubTask
				for _, subtask := range s.subtasks {
					if subtask.StoryID == story.ID {
						subtasks = append(subtasks, *subtask)
					}
				}
				story.SubTasks = subtasks
				stories = append(stories, *story)
			}
		}
		backlogCopy := *backlog
		backlogCopy.Stories = stories
		backlogs = append(backlogs, &backlogCopy)
	}

	return backlogs, nil
}

// Story operations
func (s *Service) CreateStory(req models.CreateStoryRequest) (*models.Story, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if backlog exists
	if _, exists := s.backlogs[req.BacklogID]; !exists {
		return nil, errors.New("backlog not found")
	}

	story := &models.Story{
		ID:           uuid.New().String(),
		BacklogID:    req.BacklogID,
		Title:        req.Title,
		Description:  req.Description,
		JiraURL:      req.JiraURL,
		EffortOrigin: req.EffortOrigin,
		PIC:          req.PIC,
		PlanStart:    req.PlanStart,
		PlanEnd:      req.PlanEnd,
		Status:       models.StatusTodo,
		SubTasks:     []models.SubTask{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	s.stories[story.ID] = story
	return story, nil
}

func (s *Service) GetStory(id string) (*models.Story, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	story, exists := s.stories[id]
	if !exists {
		return nil, errors.New("story not found")
	}

	// Load subtasks for this story
	var subtasks []models.SubTask
	for _, subtask := range s.subtasks {
		if subtask.StoryID == id {
			subtasks = append(subtasks, *subtask)
		}
	}
	story.SubTasks = subtasks

	return story, nil
}

func (s *Service) GetStoriesByBacklog(backlogID string) ([]*models.Story, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var stories []*models.Story
	for _, story := range s.stories {
		if story.BacklogID == backlogID {
			// Load subtasks for this story
			var subtasks []models.SubTask
			for _, subtask := range s.subtasks {
				if subtask.StoryID == story.ID {
					subtasks = append(subtasks, *subtask)
				}
			}
			storyCopy := *story
			storyCopy.SubTasks = subtasks
			stories = append(stories, &storyCopy)
		}
	}

	return stories, nil
}

// SubTask operations
func (s *Service) CreateSubTask(req models.CreateSubTaskRequest) (*models.SubTask, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if story exists
	if _, exists := s.stories[req.StoryID]; !exists {
		return nil, errors.New("story not found")
	}

	subtask := &models.SubTask{
		ID:          uuid.New().String(),
		StoryID:     req.StoryID,
		Title:       req.Title,
		Description: req.Description,
		Effort:      req.Effort,
		JiraURL:     req.JiraURL,
		PIC:         req.PIC,
		PlanStart:   req.PlanStart,
		PlanEnd:     req.PlanEnd,
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.subtasks[subtask.ID] = subtask
	return subtask, nil
}

func (s *Service) GetSubTask(id string) (*models.SubTask, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	subtask, exists := s.subtasks[id]
	if !exists {
		return nil, errors.New("subtask not found")
	}

	return subtask, nil
}

func (s *Service) GetSubTasksByStory(storyID string) ([]*models.SubTask, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var subtasks []*models.SubTask
	for _, subtask := range s.subtasks {
		if subtask.StoryID == storyID {
			subtasks = append(subtasks, subtask)
		}
	}

	return subtasks, nil
}

// Status update operations
func (s *Service) UpdateBacklogStatus(id string, status models.Status) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	backlog, exists := s.backlogs[id]
	if !exists {
		return errors.New("backlog not found")
	}

	backlog.UpdatedAt = time.Now()
	return nil
}

func (s *Service) UpdateStoryStatus(id string, status models.Status) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	story, exists := s.stories[id]
	if !exists {
		return errors.New("story not found")
	}

	story.Status = status
	story.UpdatedAt = time.Now()

	// Update actual dates based on status
	now := time.Now()
	if status == models.StatusInProgress && story.ActualStart == nil {
		story.ActualStart = &now
	} else if status == models.StatusDone && story.ActualEnd == nil {
		story.ActualEnd = &now
	}

	return nil
}

func (s *Service) UpdateSubTaskStatus(id string, status models.Status) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	subtask, exists := s.subtasks[id]
	if !exists {
		return errors.New("subtask not found")
	}

	subtask.Status = status
	subtask.UpdatedAt = time.Now()

	// Update actual dates based on status
	now := time.Now()
	if status == models.StatusInProgress && subtask.ActualStart == nil {
		subtask.ActualStart = &now
	} else if status == models.StatusDone && subtask.ActualEnd == nil {
		subtask.ActualEnd = &now
	}

	return nil
}

// Dashboard/Statistics operations
func (s *Service) GetDashboardStats() (map[string]interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	stats := map[string]interface{}{
		"total_backlogs": len(s.backlogs),
		"total_stories":  len(s.stories),
		"total_subtasks": len(s.subtasks),
	}

	// Count by status
	storyStatusCount := make(map[models.Status]int)
	subtaskStatusCount := make(map[models.Status]int)

	for _, story := range s.stories {
		storyStatusCount[story.Status]++
	}

	for _, subtask := range s.subtasks {
		subtaskStatusCount[subtask.Status]++
	}

	stats["story_status"] = storyStatusCount
	stats["subtask_status"] = subtaskStatusCount

	return stats, nil
}
