package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/marcy-ot/go-api-server-by-build-in/domain"
	"github.com/marcy-ot/go-api-server-by-build-in/repository"
)

// todoHandler is handler for search todo
func todoHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: filter by title from query parameter
	title := r.URL.Query().Get("title")
	je := json.NewEncoder(w)
	je.SetIndent("", "  ")
	je.Encode(repository.SearchTodo(domain.TodoSearchCondition{Title: title}))
}

// todoDetailHandler is handler for get todo detail
func todoDetailHandler(w http.ResponseWriter, r *http.Request) {
	requiredID := r.PathValue("ID")
	id, _ := strconv.Atoi(requiredID)
	todo, err := repository.FindTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	je := json.NewEncoder(w)
	je.SetIndent("", "  ")
	je.Encode(todo)
}

// todoCreateHandler is handler for create todo
func todoCreateHandler(w http.ResponseWriter, r *http.Request) {
	// Validate http Method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Validate request body
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoReq := domain.Todo{
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create todo
	if err := repository.ResisterTodo(todoReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	je := json.NewEncoder(w)
	je.SetIndent("", "  ")
	je.Encode(true)
}

func main() {
	http.HandleFunc("/api/v1/todo", todoHandler)
	http.HandleFunc("/api/v1/todo/{ID}", todoDetailHandler)
	http.HandleFunc("/api/v1/todo/create", todoCreateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
