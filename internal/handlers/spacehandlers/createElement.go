package spacehandlers

import "net/http"


func CreateElementHandler (w http.ResponseWriter, r * http.Request) {

	w.Header().Set("Content-Type","application/json")
}

