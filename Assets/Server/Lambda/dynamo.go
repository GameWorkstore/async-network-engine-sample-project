package main

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Global Vars
var dynamoConnection *dynamodb.DynamoDB
var tableName = "lks-users"
var tableAttName = "name"
var tableAttID = "id"
var tableAttCoins = "coins"

func getCoins(mapped map[string]*dynamodb.AttributeValue) int {
	v, err := strconv.Atoi(*mapped[tableAttCoins].N)
	if err != nil {
		panic("conversion error")
	}
	return v
}

func getDynamoConnection() *dynamodb.DynamoDB {
	if dynamoConnection == nil {
		session := session.Must(session.NewSession())
		dynamoConnection = dynamodb.New(session)
	}
	return dynamoConnection
}
