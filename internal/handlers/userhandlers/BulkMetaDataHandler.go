package userHandler

import (
	"fmt"
	"net/http"
)

func MetaDataHandler (w  http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"hello from metadata handler")
}