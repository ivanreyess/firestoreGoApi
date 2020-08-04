package service

import (
	"../entity"
	"../repository"
	"errors"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

var (
	repo repository.PostRepository
)

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the title is empty")
		return err
	}
	if post.Text == "" {
		err := errors.New("the text is empty")
		return err
	}
	return nil
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
