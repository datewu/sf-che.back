package api

import (
	"net/http"

	"github.com/datewu/gtea/handler"
)

func hw(w http.ResponseWriter, r *http.Request) {
	msg := map[string]interface{}{
		"message": "Hello, World!",
		"you":     r.RemoteAddr,
	}
	handler.OKJSON(w, msg)
}
