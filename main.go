package main

import (
	"fmt"
	"os"
	"strconv"

	"todo-cli/internal/storage"
	"todo-cli/internal/tasks"
)

const filename = "tasks.json"

func main() {
	task, err := storage.LoadTasks(filename)
	if err != nil {
		fmt.Println("Erro ao carregar tarefas:", err)
		return
	}

	service := tasks.NewService(task)

	if len(os.Args) < 2 {
		fmt.Println("Comandos: add, list, done, remove")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Uso: go run . add \"nome da tarefa\"")
			return
		}

		newTask := service.Add(os.Args[2])
		err = storage.SaveTasks(filename, service.Tasks())
		if err != nil {
			fmt.Println("Erro ao salvar tarefas:", err)
			return
		}

		fmt.Println("Tarefa criada:", newTask.Title)

	case "list":
		if len(service.List()) == 0 {
			fmt.Println("Lista de tarefas vazia")
		} else {
			for _, task := range service.List() {
				status := "pendente"
				if task.Done {
					status = "concluída"
				}

				fmt.Printf("[%d] %s - %s\n", task.Id, task.Title, status)
			}
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Uso: go run . done 1")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID inválido")
			return
		}

		err = service.Done(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = storage.SaveTasks(filename, service.Tasks())
		if err != nil {
			fmt.Println("Erro ao salvar tarefas:", err)
			return
		}

		fmt.Println("Tarefa concluída")

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Uso: go run . remove 1")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID inválido")
			return
		}

		err = service.Remove(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = storage.SaveTasks(filename, service.Tasks())
		if err != nil {
			fmt.Println("Erro ao salvar tarefas:", err)
			return
		}

		fmt.Println("Tarefa removida")

	default:
		fmt.Println("Comando inválido")
	}
}
