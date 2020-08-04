package repository

import (
	"../entity"
)

//PostRepository defines an interface
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
