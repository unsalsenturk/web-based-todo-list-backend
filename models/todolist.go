package models

type Todo struct {
	ID          uint   `json:"id" `
	Description string `json:"description"`
}

type DataResponse map[string]Todo
