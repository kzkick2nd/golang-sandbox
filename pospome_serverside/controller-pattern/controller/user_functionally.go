package controller

import (
	"fmt"
	"net/http"
)

// 関数ごとに切り分けても見通しには限界がある、小規模向け
func Handler(w http.ResponseWriter, r *http.Request) {
	n := r.FormValuse("name")
	if n == "" {
		fmt.Fprint(w, "name error")
		return
	}

	if err := InsertUser(n); err != nil {
		fmt.Fprint(w, "name error")
	}

	fmt.Fprint(w, "success")
}

func InsertUser(name string) error {
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
