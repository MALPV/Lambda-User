package secret

import (
	"encoding/json"
	"fmt"

	"github.com/MALPV/Lambda-User/awsgo"
	"github.com/MALPV/Lambda-User/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secret string) (models.SecretRDSJson, error) {

	var dataSecret models.SecretRDSJson

	fmt.Println("getSecret -> Init " + secret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	secretValue, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secret),
	})

	if err != nil {
		fmt.Println("error: GetSecret " + err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*secretValue.SecretString), &dataSecret)

	fmt.Println("getSecret -> Success " + secret)

	return dataSecret, nil

}
