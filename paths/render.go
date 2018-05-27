package path

import "net/http"

type render struct{}

func Render(json []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
