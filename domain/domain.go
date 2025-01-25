package domain

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TodoSearchCondition is condition for search todo
type TodoSearchCondition struct {
	Title string
}
