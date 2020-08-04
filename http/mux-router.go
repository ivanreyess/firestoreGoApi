package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouter struct {
}

func NewMuxRouter() Router {
	return &MuxRouter{}
}

var muxDispatcher = mux.NewRouter()

func (m *MuxRouter) GET(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (m *MuxRouter) POST(uri string, f func(rw http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (m *MuxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP Server running on port: %v", port)
	_ = http.ListenAndServe(port, muxDispatcher)
}
