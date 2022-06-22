// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Task struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
	Views     int    `json:"views"`
}

type TaskInput struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
	View      int    `json:"view"`
}
