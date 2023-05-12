package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/MALPV/Lambda-User/awsgo"
	"github.com/MALPV/Lambda-User/bd"
	"github.com/MALPV/Lambda-User/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(InvokeLambda)
}

func InvokeLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InitAWS()

	if !ValidateParameters() {
		err := errors.New("error: must send SecretName")
		fmt.Println(err)
		return event, err
	}

	var user models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			user.UserEmail = att
			fmt.Println("userAttributes -> Email: ", user.UserEmail)
		case "sub":
			user.UserUUID = att
			fmt.Println("userAttributes -> Sub: ", user.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("error: Reading Secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(user)
	return event, err

}

func ValidateParameters() bool {

	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam

}
