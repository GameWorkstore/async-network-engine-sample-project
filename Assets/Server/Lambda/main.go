package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ImplementationFailure handles requests from API Gateway when function wasn't implemented properly
func ImplementationFailure(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var resp = events.APIGatewayProxyResponse{}
	resp.Body = "Implementation failure, verify if main contains the name of function"
	resp.StatusCode = 200

	return resp, nil
}

// Request is struct
type Request struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetData is a function
func GetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("GetData")

	var rqt Request
	err := json.Unmarshal([]byte(request.Body), &rqt)
	if err != nil {
		var fail = events.APIGatewayProxyResponse{}
		fail.Body = "Unmashal error"
		fail.StatusCode = 200
		return fail, nil
	}

	var resp = events.APIGatewayProxyResponse{}
	resp.Body = "Welcome, " + rqt.Name
	resp.StatusCode = 200

	return resp, nil
}

// SetData is a function
func SetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("SetData")

	var rqt Request
	var err error
	err = json.Unmarshal([]byte(request.Body), &rqt)
	if err != nil {
		var fail = events.APIGatewayProxyResponse{}
		fail.Body = "Unmashal error"
		fail.StatusCode = 200
		return fail, nil
	}

	var idAttrib dynamodb.AttributeValue
	idAttrib.S = &rqt.ID

	var nameAttrib dynamodb.AttributeValue
	nameAttrib.S = &rqt.Name

	var attributes map[string]*dynamodb.AttributeValue
	attributes["id"] = &idAttrib
	attributes["name"] = &nameAttrib

	tableName := "lks-users"

	var put dynamodb.PutItemInput
	put.TableName = &tableName
	put.Item = attributes

	_, err = getDynamoConnection().PutItem(&put)
	if err != nil {
		var fail = events.APIGatewayProxyResponse{}
		fail.Body = "Unmashal error"
		fail.StatusCode = 200
		return fail, nil
	}

	var resp = events.APIGatewayProxyResponse{}
	resp.Body = "Updated, " + rqt.Name
	resp.StatusCode = 200

	return resp, nil
}

func main() {

	var functionName = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	fmt.Println("FunctionName" + functionName)
	switch functionName {
	case "lkssetdata":
		lambda.Start(SetData)
		return
	case "lksgetdata":
		lambda.Start(GetData)
		return
	default:
		lambda.Start(ImplementationFailure)
		return
	}
}
