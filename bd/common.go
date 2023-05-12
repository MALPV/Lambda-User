package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/MALPV/Lambda-User/models"
	"github.com/MALPV/Lambda-User/secret"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {

	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err

}

func ConnectDb() error {

	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("error: ConnectDb " + err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("error: ConnectDb " + err.Error())
		return err
	}

	fmt.Println("connectDb -> Success")
	return nil

}

func ConnStr(key models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string

	dbUser = key.Username
	authToken = key.Password
	dbEndPoint = key.Host
	dbName = "malpv"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	fmt.Println("ConnStr -> Success")
	return dsn

}
