package task

type Service interface {
	GetTasks() ([]GetTasksResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTasks() ([]GetTasksResponse, error) {
	return s.repository.GetTasks()
}
