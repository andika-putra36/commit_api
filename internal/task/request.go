package task

type CreateTaskRequest struct {
	Title    string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle" binding:"required"`
	StartAt  string `json:"start_at" binding:"required"`
	FinishAt string `json:"finish_at" binding:"required"`
}

type UpdateTaskRequest struct {
	Title    string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle" binding:"required"`
	StartAt  string `json:"start_at" binding:"required"`
	FinishAt string `json:"finish_at" binding:"required"`
	IsDone   *bool  `json:"is_done" binding:"required"`
}
