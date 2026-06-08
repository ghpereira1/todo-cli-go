package storage

import (
	"encoding/json"
	"os"

	"todo-cli/internal/tasks"
)

func LoadTasks(filename string) ([]tasks.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []tasks.Task{}, nil
		}

		return nil, err
	}

	var tasks []tasks.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(filename string, tasks []tasks.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
