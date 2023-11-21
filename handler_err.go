package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	ResponseWithJSON(w, 400, "Something went wrong")
}