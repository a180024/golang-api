package db

import (
	"github.com/a180024/golang-api/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Init() *dynamodb.DynamoDB {
	c := config.GetConfig()
	sess := session.New(&aws.Config{
		Region:      aws.String(c.GetString("dbRegion")),
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    aws.String(c.GetString("dbEndpointURL")),
		DisableSSL:  aws.Bool(true),
	})

	// Return dynamodb client
	return dynamodb.New(sess)
}
