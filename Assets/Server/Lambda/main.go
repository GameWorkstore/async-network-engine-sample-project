package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

// RequestGet request input for get function
type RequestGet struct {
	ID string `json:"id"`
}

// RequestUpdate request input for update function
type RequestUpdate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetData is a function
func GetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("GetData")

	var rqt RequestGet
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

	var get dynamodb.GetItemInput
	get.TableName = &tableName
	get.AttributesToGet = []*string{&tableAttName}
	get.Key = make(map[string]*dynamodb.AttributeValue)
	get.Key[tableAttID] = &idAttrib

	output, err := getDynamoConnection().GetItem(&get)
	if err != nil {
		var fail = events.APIGatewayProxyResponse{}
		fail.Body = "Get error:" + err.Error()
		fail.StatusCode = 200
		return fail, nil
	}

	var resp = events.APIGatewayProxyResponse{}
	resp.Body = "Welcome back, " + *output.Item[tableAttName].S
	resp.StatusCode = 200

	return resp, nil
}

// SetData is a function
func SetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("SetData")

	var rqt RequestUpdate
	var err error
	err = json.Unmarshal([]byte(request.Body), &rqt)
	if err != nil {
		var fail = events.APIGatewayProxyResponse{}
		fail.Body = "Unmashal error:" + err.Error()
		fail.StatusCode = 200
		return fail, nil
	}

	var idAttrib dynamodb.AttributeValue
	idAttrib.S = &rqt.ID

	var nameAttrib dynamodb.AttributeValue
	nameAttrib.S = &rqt.Name

	var sint = strconv.Itoa(403)
	var coinsAttrib dynamodb.AttributeValue
	coinsAttrib.N = &sint

	var put dynamodb.PutItemInput
	put.TableName = &tableName
	put.Item = make(map[string]*dynamodb.AttributeValue)
	put.Item[tableAttID] = &idAttrib
	put.Item[tableAttName] = &nameAttrib
	put.Item[tableAttCoins] = &coinsAttrib

	_, err = getDynamoConnection().PutItem(&put)
	if err != nil {
		var fail = events.APIGatewayProxyResponse{}
		fail.Body = "Set error:" + err.Error()
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
