package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func AWSInit() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Unable to load SDK config, " + err.Error())
	}
}