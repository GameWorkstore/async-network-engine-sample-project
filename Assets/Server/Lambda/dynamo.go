package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Global Vars
var dynamoConnection *dynamodb.DynamoDB
var tableName = "lks-users"
var tableAttName = "name"
var tableAttID = "id"
var tableAttCoins = "coins"

func getDynamoConnection() *dynamodb.DynamoDB {
	if dynamoConnection == nil {
		session := session.Must(session.NewSession())
		dynamoConnection = dynamodb.New(session)
	}
	return dynamoConnection
}
