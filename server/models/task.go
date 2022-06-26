package models

type Task struct {
	ID        string `json:"id"`
	Views     int    `json:"views"`
	Timestamp string `json:"timestamp"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
