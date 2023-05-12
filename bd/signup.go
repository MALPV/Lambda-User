package bd

import (
	"fmt"

	"github.com/MALPV/Lambda-User/models"
	"github.com/MALPV/Lambda-User/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sign models.SignUp) error {

	fmt.Println("signUp -> Init")

	err := ConnectDb()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sign.UserEmail + "','" + sign.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println("Executing query: ", query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println("error: Execute Query " + err.Error())
		return err
	}

	fmt.Println("signUp -> Success")

	return nil

}
