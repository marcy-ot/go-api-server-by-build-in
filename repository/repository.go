package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/marcy-ot/go-api-server-by-build-in/domain"
)

var todoFilePath = "data_store/todo.json"

// ReadTodo is read todo.json and return Todo list
func ReadTodo() []domain.Todo {
	data, err := os.ReadFile(todoFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var todoList []domain.Todo
	if err := json.Unmarshal(data, &todoList); err != nil {
		log.Fatal(err)
	}
	return todoList
}

func SearchTodo(condition domain.TodoSearchCondition) []domain.Todo {
	todoList := ReadTodo()
	var result []domain.Todo
	// filter by condition
	for _, todo := range todoList {
		if strings.Contains(todo.Title, condition.Title) {
			result = append(result, todo)
		}
	}
	return result
}

// FindTodo is find todo by id
func FindTodo(id int) (domain.Todo, error) {
	todoList := ReadTodo()
	for _, todo := range todoList {
		if todo.Id == id {
			return todo, nil
		}
	}
	return domain.Todo{}, fmt.Errorf("todo not found")
}

// resisterTodo is resister todo
func ResisterTodo(todo domain.Todo) error {
	todoList := ReadTodo()

	var id int
	for _, t := range todoList {
		if id < t.Id {
			id = t.Id
		}
	}
	todo.Id = id + 1

	todoList = append(todoList, todo)
	json, err := json.Marshal(todoList)

	if err != nil {
		return err
	}
	if err := os.WriteFile(todoFilePath, json, 0644); err != nil {
		return err
	}
	return nil
}
