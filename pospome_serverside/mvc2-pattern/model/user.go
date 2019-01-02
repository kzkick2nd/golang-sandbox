package model

type User struct {
	ID     int
	Name   int
	Status string
}

func (u *User) IsAdmin() bool {
	return u.Status == "ADMIN"
}

// func UpdateUser(conn *db.Conn, user *User) error {
// 	sql := `Update users SET name = "%s", status = "%s" WHERE id = %d`
// 	q := query.Create(sql, user.Name, user.Status, user.ID)
// 	return db.Exec(conn, q)
// }

func (u *User) Update() error {
	sql := `Update users SET name = "%s", status = "%s" WHERE id = %d`
	q := query.Create(sql, u.Name, u.Status, u.ID)
	return db.Exec(conn, q)
}
