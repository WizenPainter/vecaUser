package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/WizenPainter/vecaUser/models"
	"github.com/WizenPainter/vecaUser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Beginning Register")

	err := DbConnect()

	if err != nil {
		return err
	}
	defer Db.Close()

	sentence := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('"+sig.UserEmail+"', '"+sig.UserUUID+"', '"+tools.DateMySQL()+"')"
	fmt.Println("Executing SQL code: " + sentence)
	_, err = Db.Exec(sentence)
	
	if err != nil {
		fmt.Println("Error executing SQL command: ", err.Error())
		return err
	}
	fmt.Println("SignUp > Successfull")
	return nil
}