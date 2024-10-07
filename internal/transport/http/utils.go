package restTransport

import (
	"fmt"
	"net/http"
)

func errorHandler(e error, w http.ResponseWriter) {
	fmt.Println(e)
	w.WriteHeader(400)
	w.Write([]byte(e.Error()))
}
