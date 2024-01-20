package models

import "time"

type Book struct {
	Name        string     `json:"book"`
	Genre       string     `json:"genre"`
	DateRelease *time.Time `json:"date_release"`
	AuthorID    string     `json:"author_id"`
}
