package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/WizenPainter/vecaUser/models"
	"github.com/WizenPainter/vecaUser/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("Error connecting to db, Error Message: " + err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Error pinging db, Error Message: " + err.Error())
		return err
	}

	fmt.Println("Successfull connection to DB")
	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "vecastore"
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println("DNS set to: " + dns)
	return dns
}