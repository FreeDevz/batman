package main

import (
	"fmt"
	"app/model"
)

var currentId int

var todos model.Todos

// Give us some seed data
func init() {
	RepoCreateTodo(model.Todo{Name: "Write presentation", Completed: true})
	RepoCreateTodo(model.Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) model.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return model.Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t model.Todo) model.Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
