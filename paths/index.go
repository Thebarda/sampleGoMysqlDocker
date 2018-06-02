package path

import (
	"encoding/json"
	"net/http"

	"../requests"
)

func User(w http.ResponseWriter, r *http.Request) {
	users := requests.GetAllUsers()
	jsonResult, _ := json.Marshal(users)
	Render(jsonResult, w)
}
