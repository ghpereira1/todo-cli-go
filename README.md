# Todo CLI em Go

CLI simples de gerenciamento de tarefas desenvolvida em Go.

## Funcionalidades

- Adicionar tarefas
- Listar tarefas
- Marcar tarefas como concluídas
- Remover tarefas
- Persistir dados em JSON
- Testes unitários com go test

## Tecnologias

- Go
- JSON
- Go Test

## Como rodar

```bash
go run . add "Estudar Go"
go run . list
go run . done 1
go run . remove 1
