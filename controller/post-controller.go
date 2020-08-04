package controller

import (
	"../entity"
	"../errors"
	"../service"
	"encoding/json"
	"net/http"
)

type PostController interface {
	GetPosts(rw http.ResponseWriter, r *http.Request)
	AddPost(rw http.ResponseWriter, r *http.Request)
}

type controller struct{}

var (
	s service.PostService
)

func NewPostController(service service.PostService) PostController {
	s = service
	return &controller{}
}

func (c *controller) GetPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	posts, err := s.FindAll()
	response, err := json.Marshal(posts)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error converting posts array"})
		return
	}
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(response)
}

func (c *controller) AddPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := s.Validate(&post)
	if err1 != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Invalid data"})
		return
	}
	_, err2 := s.Create(&post)
	if err2 != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Could not save post"})
		return
	}
	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(post)

}
