package controller

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	n := r.FormValuse("name")
	if n == "" {
		fmt.Fprint(w, "name error")
		return
	}

	conn := db.NewConnection()
	conn.BeginTransaction()

	u := &User{
		Name: n,
	}

	if err := db.Insert(conn, u); err != nil {
		db.Rollback()
		fmt.Fprint(w, "db error")
		return
	}
	conn.Commit()

	fmt.Fprint(w, "success")
}

type User struct {
	Name string
}
