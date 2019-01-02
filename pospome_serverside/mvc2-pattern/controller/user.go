package controller

import (
	"net/http"
	"strconv"
)

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	u := &model.User{}

// 	b, err := json.Marchal(u)
// 	if err != nil {
// 		fmt.Fprint(w, err.Error())
// 		return
// 	}
// 	fmt.Fprint(w, string(b))
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValuse("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	conn := model.NewDBConn()
	user := model.GetUserByID(conn, id)
	user.UpdateAdminStatus()
	if err := model.UpdateUser(conn, user); err != nil {
		conn.Rollback()
		return
	}
	conn.Commit()
	view.RenderUserPage(w, user)
}
