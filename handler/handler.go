package handler

import (
	"net/http"
)

func GetHealth(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world!!"))
}

func PostMessageCountRank(res http.ResponseWriter, req *http.Request) {

}

func PostMessageContentCountRank(res http.ResponseWriter, req *http.Request) {

}
