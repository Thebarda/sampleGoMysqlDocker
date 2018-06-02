package path

import (
	"encoding/json"
	"net/http"

	"../models"
	"../requests"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	res := requests.Insert(user)
	jsonResult, _ := json.Marshal(res)
	Render(jsonResult, w)
}
