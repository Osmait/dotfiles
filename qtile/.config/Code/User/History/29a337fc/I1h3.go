package domain

import "github.com/google/uuid"

type Post struct {
	Id        uuid.UUID
	Title     string
	CreatedAt string
}
