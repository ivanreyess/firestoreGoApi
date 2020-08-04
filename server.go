package main

import (
	"./controller"
	customHttp "./http"
	"./repository"
	"./service"
)

var (
	r              = repository.NewFirestoreRepository()
	s              = service.NewPostService(r)
	router         = customHttp.NewMuxRouter()
	postController = controller.NewPostController(s)
)

func main() {

	const port string = ":8000"
	router.GET("/posts", postController.GetPosts)
	router.POST("/posts", postController.AddPost)
	router.SERVE(port)
}
