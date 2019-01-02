package controller

import "net/http"

// サービス無し初期コントローラー
// func AddProjectAndTask(w http.ResponseWriter, r *http.Request) {
// 	projectName := getProjectNameFrom(r)
// 	taskName := getTaskNameFrom(r)

// 	conn := model.NewDBConn()

// 	p, err := model.NewProject(projectName)
// 	if err != nil {
// 		panic("error")
// 	}

// 	if err := model.AddProject(conn, p); err != nil {
// 		conn.Rollback()
// 		panic("error")
// 	}

// 	t, err := model.NewTask(taskName)
// 	if err != nil {
// 		pacnic("error")
// 	}

// 	if err := model.AddTask(t); err != nil {
// 		conn.Rollback()
// 		panic("error")
// 	}

// 	conn.Commit
// }

// 複数レイヤー式で書き換え
func AddProjectAndTask(w http.ResponseWriter, r *http.Request) {
	projectName := getProjectNameFrom(r)
	taskName := getTaskNameFrom(r)

	// 複数サービスは複数サービスレイヤーにまたがるので、トランザクションはコントローラー側でもつ
	conn := model.NewDBConn()

	if err := service.AddProject(conn, projectName); err != nil {
		conn.Rollback()
		panic("error")
	}

	if err := service.AddTask(conn, taskName); err != nil {
		con.Rollback()
		panic("error")
	}

	conn.Commit
}

// サービスレイヤーに移動する前
func AddTask(w http.ResponseWriter, r *http.Request) {
	projectID := getProjectIDFrom(r)
	taskName := getTaskNameFrom(r)
	assingneeID := getTaskAssighneeIDFrom(r)

	conn := model.NewDBConn()
	p, err := model.GetProjectByID(conn, projectID)
	if err != nil {
		panic("error")
	}
	if !p.IsAvailable() {
		paic("error")
	}

	if yes, err := model.DoesExistSameNameTask(conn, taskName); err != nil || !yes {
		panic("error")
	}

	u, err := model.GetUserByID(con, assingneeID)
	if err != nil {
		panic("error")
	}
	if !u.IsDeveloper() {
		panic("error")
	}

	if err := model.AddTask(conn, taskName); err != nil {
		panic("error")
	}
}

// サービスレイヤーに移動1。モデル毎に移動。
func AddTask(w http.ResponseWriter, r *http.Request) {
	projectID := getProjectIDFrom(r)
	taskName := getTaskNameFrom(r)
	assingneeID := getTaskAssighneeIDFrom(r)

	conn := model.NewDBConn()

	if err := service.IsAvailableProject(conn, projectID); err != nil {
		panic("error")
	}

	if err := service.IsDeveloperUser(conn, assigneeID); err != nil {
		panic("error")
	}

	if err := service.AddTask(conn, taskName); err != nil {
		panic("error")
	}
}

// サービスレイヤーに移動2。他で利用するコントローラーをサービスに切り出し（抽象化サービス的な？）
func AddTask(w http.ResponseWriter, r *http.Request) {
	projectID := getProjectIDFrom(r)
	taskName := getTaskNameFrom(r)
	assingneeID := getTaskAssighneeIDFrom(r)

	conn := model.NewDBConn()

	// 別メソッドでも利用するから抽象化サービスへ切り出し
	if err := service.IsAvailableProject(conn, projectID); err != nil {
		panic("error")
	}

	if yes, err := model.DoesExistSameNameTask(conn, taskName); err != nil || !yes {
		panic("error")
	}

	u, err := model.GetUserByID(con, assingneeID)
	if err != nil {
		panic("error")
	}
	if !u.IsDeveloper() {
		panic("error")
	}

	if err := model.AddTask(conn, taskName); err != nil {
		panic("error")
	}
}
func AddMessage(w http.ResponseWriter, r *http.Request) {
	projectID := getProjectIDFrom(r)
	conn := model.NewDBConn()
	if err := service.IsAvailableProject(conn, projectID); err != nil {
		panic("error")
	}
	// 以下略
}

// サービスレイヤーに移動3。まとめてバリデーションとして切り出し。
func AddTask(w http.ResponseWriter, r *http.Request) {
	projectID := getProjectIDFrom(r)
	taskName := getTaskNameFrom(r)
	assingneeID := getTaskAssighneeIDFrom(r)

	conn := model.NewDBConn()

	if err := service.CanAddTask(conn, projectID, assigneeID, taskName); err != nil {
		panic("error")
	}

	if err := model.AddTask(taskName); err != nil {
		panic("error")
	}
}
