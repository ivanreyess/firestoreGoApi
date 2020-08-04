package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity"
	"./repository"
)

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func getPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	response, err := json.Marshal(posts)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`"error":"Error converting posts array"`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(response)
}

func addPost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(post)

}
