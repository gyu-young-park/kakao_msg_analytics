package customrouter

import "github.com/gorilla/mux"

type CustomRouter interface {
	SetRouteWithHandler(*mux.Router)
}
