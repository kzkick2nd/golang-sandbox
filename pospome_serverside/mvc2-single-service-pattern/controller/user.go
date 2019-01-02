package controller

import (
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 入力値をモデルに適した値に加工する
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	// // モデルを利用して処理をすすめる
	// // NOTE: モデルを操作しているロジック = サービスレイヤーに移動
	// conn := model.NewDBConn()
	// user := model.GetUserByID(conn, id)
	// user.UpdateAdminStatus()
	// if err := model.UpdateUser(conn, user); err != nil {
	// 	conn.Rollback()
	// 	return
	// }
	// conn.Commit()

	// service layerのメソッドを呼び出すだけ
	if err := service.UpdateUserStatus(id); err != nil {
		return
	}

	// 出力はビューを利用する
	view.RenderUserPage(w, user)
}

func UploadBackgrounbImage(w http.ResponseWriter, r *http.Request) {
	img := getImageFrom(r)
	service.UploadBackgroundImage(img)
	return
}

func UploadUserImage(w http.ResponseWriter, r *http.Request) {
	img := getImageFrom(r)
	service.UploadBackgroundImage(img)
	return
}
