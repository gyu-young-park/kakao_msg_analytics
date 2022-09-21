package customrouter

import (
	"gyu/handler"

	"github.com/gorilla/mux"
)

type MesageRouter_v1 struct {
}

func (m *MesageRouter_v1) SetRouteWithHandler(router *mux.Router) {
	router.HandleFunc("/health", handler.GetHealth).Methods("GET")
	router.HandleFunc("/message/rank", handler.PostMessageCountRank).Methods("POST")
	router.HandleFunc("/message/content/rank/{limit:[0-9]+}", handler.PostMessageContentCountRank).Methods("POST")
}
