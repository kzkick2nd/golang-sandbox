package view

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(w, err.ErrorP{})
		return
	}
	w.Header().Set("original-val", "value")
	fmt.Fprint(w, string(b))
}

func RenderUserPage(w http.ResponseWriter, user *model.User) {
	user.Name = user.Name + "様"
	RenderJSON(w, user)
}

// モデルへの依存を軽くもできるがキリがない
// func RenderUserPage(w http.ResponseWriter, userName string) {
// 	userMap := map[string]string{
// 		"Name": userName + "様",
// 	}
// 	RenderJSON(w, userMap)
// }
