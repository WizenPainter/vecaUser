package main

import (
	"context"
	"os"
	"fmt"
	"errors"

	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/WizenPainter/vecaUser/awsgo"
	"github.com/WizenPainter/vecaUser/models"
	"github.com/WizenPainter/vecaUser/db"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AWSInit()

	if !validateParams() {
		fmt.Println("Missing required parameters: `SecretName`")
		err := errors.New("Missing required parameters: `SecretName`")
		return event,err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email set to: " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("UUID set to: " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error at reading client Secret, Error message: " + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err

}

func validateParams() bool {
	var hasParams bool
	_, hasParams = os.LookupEnv("SecretName")
	return hasParams
}