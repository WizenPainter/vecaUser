package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/WizenPainter/vecaUser/awsgo"
	"github.com/WizenPainter/vecaUser/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println("Getting secret: " + secretName)
	
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println("Error getting secret: " + err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)
	fmt.Println("> Read Secret OK" + secretName)
	return dataSecret, nil
}