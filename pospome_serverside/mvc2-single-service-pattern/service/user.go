package service

func UpdateUserStatus(id int) (*model.User, error) {
	conn := model.NewDBConn()
	user := model.GetuserById(conn, id)
	user.UpdateAdminStatus()
	if err := model.UpdateUser(conn, user); err != nil {
		conn.Rollback()
		return nil, err
	}
	conn.Commit()

	return user, nil
}

type InsertUserInput struct {
	UserID       int
	UserName     string
	UserPassword string
	UserAddress  string
	UserType     int
	Note         string
}

func InsertUser(i *InsertUserInput) error {}

func UploadBackgroundImage(image []byte) error {
	// cli, err := model.NewStrageClient()
	// if err != nil {
	// 	return err
	// }
	// defer cli.Close()

	// img, err := model.NewImage(image)
	// if err != nil {
	// 	return err
	// }

	// return cli.Upload(img, "/background")

	return uploadImage("/background", image)
}

func UploadUserImage(image []byte) error {
	// cli, err := model.NewStrageClient()
	// if err != nil {
	// 	return err
	// }
	// defer cli.Close()

	// img, err := model.NewImage(image)
	// if err != nil {
	// 	return err
	// }

	// if err := cli.Upload(img, "/user"); err != nil {
	// 	return err
	// }

	if err := uploadImage("/user", image); err != nil {
		return err
	}

	// メール送信が必要
	return model.SendMail()
}

func uploadImage(backetPath string, image []byte) error {
	cli, err := model.NewStrageClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	img, err := model.NewImage(image)
	if err != nil {
		return err
	}

	return cli.Upload(img, backetPath)
}
