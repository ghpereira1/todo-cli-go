package tasks

import "testing"

func TestServiceAdd(t *testing.T) {
	service := NewService([]Task{})

	newTask := service.Add("Estudar Go")

	if newTask.Id != 1 {
		t.Errorf("esperado Id 1, recebeu %d", newTask.Id)
	}

	if newTask.Title != "Estudar Go" {
		t.Errorf("esperado título 'Estudar Go', recebeu %s", newTask.Title)
	}

	if newTask.Done {
		t.Error("esperado Done false, recebeu true")
	}

	if newTask.CreatedAt.IsZero() {
		t.Error("esperado CreatedAt preenchido, recebeu valor zero")
	}

	if len(service.Tasks()) != 1 {
		t.Errorf("esperado 1 tarefa no service, recebeu %d", len(service.Tasks()))
	}
}

func TestServiceAddWithExistingTasks(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Tarefa 1", Done: false},
		{Id: 2, Title: "Tarefa 2", Done: false},
	}

	service := NewService(tasks)

	newTask := service.Add("Tarefa 3")

	if newTask.Id != 3 {
		t.Errorf("esperado Id 3, recebeu %d", newTask.Id)
	}
}

func TestServiceList(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Estudar Go", Done: false},
		{Id: 2, Title: "Estudar Docker", Done: true},
	}

	service := NewService(tasks)

	result := service.List()

	if len(result) != 2 {
		t.Errorf("esperado 2 tarefas, recebeu %d", len(result))
	}

	if result[0].Title != "Estudar Go" {
		t.Errorf("esperado 'Estudar Go', recebeu %s", result[0].Title)
	}

	if result[1].Title != "Estudar Docker" {
		t.Errorf("esperado 'Estudar Docker', recebeu %s", result[1].Title)
	}
}

func TestServiceDone(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Estudar Go", Done: false},
		{Id: 2, Title: "Estudar Docker", Done: false},
	}

	service := NewService(tasks)

	err := service.Done(1)

	if err != nil {
		t.Errorf("não esperava erro, recebeu %v", err)
	}

	result := service.Tasks()

	if !result[0].Done {
		t.Error("esperado tarefa com Id 1 concluída")
	}

	if result[1].Done {
		t.Error("não esperava tarefa com Id 2 concluída")
	}
}

func TestServiceDoneTaskNotFound(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Estudar Go", Done: false},
	}

	service := NewService(tasks)

	err := service.Done(99)

	if err == nil {
		t.Error("esperava erro para tarefa inexistente, recebeu nil")
	}
}

func TestServiceRemove(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Estudar Go", Done: false},
		{Id: 2, Title: "Estudar Docker", Done: false},
	}

	service := NewService(tasks)

	err := service.Remove(1)

	if err != nil {
		t.Errorf("não esperava erro, recebeu %v", err)
	}

	result := service.Tasks()

	if len(result) != 1 {
		t.Errorf("esperado 1 tarefa após remoção, recebeu %d", len(result))
	}

	if result[0].Id != 2 {
		t.Errorf("esperado que sobrasse tarefa Id 2, recebeu Id %d", result[0].Id)
	}
}

func TestServiceRemoveTaskNotFound(t *testing.T) {
	tasks := []Task{
		{Id: 1, Title: "Estudar Go", Done: false},
	}

	service := NewService(tasks)

	err := service.Remove(99)

	if err == nil {
		t.Error("esperava erro para tarefa inexistente, recebeu nil")
	}

	if len(service.Tasks()) != 1 {
		t.Errorf("esperado manter 1 tarefa, recebeu %d", len(service.Tasks()))
	}
}
