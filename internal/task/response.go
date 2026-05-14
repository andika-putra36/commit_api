package task

import "time"

type GetTasksResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	StartAt   string    `json:"start_at"`
	FinishAt  string    `json:"finish_at"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
