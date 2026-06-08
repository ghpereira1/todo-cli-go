package tasks

import (
	"errors"
	"time"
)

type Service struct {
	tasks []Task
}

func NewService(tasks []Task) *Service {
	return &Service{tasks: tasks}
}

func (s *Service) Add(title string) Task {
	task := Task{
		Id:        s.nextID(),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	s.tasks = append(s.tasks, task)

	return task
}

func (s *Service) List() []Task {
	return s.tasks
}

func (s *Service) Done(id int) error {
	for i := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks[i].Done = true
			return nil
		}
	}

	return errors.New("tarefa não encontrada")
}

func (s *Service) Remove(id int) error {
	for i := range s.tasks {
		if s.tasks[i].Id == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("tarefa não encontrada")
}

func (s *Service) Tasks() []Task {
	return s.tasks
}

func (s *Service) nextID() int {
	maxID := 0

	for _, task := range s.tasks {
		if task.Id > maxID {
			maxID = task.Id
		}
	}

	return maxID + 1
}
