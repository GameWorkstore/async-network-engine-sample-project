package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	ase "github.com/GameWorkstore/async-network-engine-go"
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

// GetData is a function
func GetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("GetData:1")

	// Cross-Origin Domain - WebGL
	respOptions, shouldStop := ase.AWSPreFlight(request)
	if shouldStop {
		return respOptions, nil
	}

	var rqt = GetUserRequest{}
	err := ase.AWSDecode(request, &rqt)
	if err != nil {
		return ase.AWSError(ase.Transmission_MarshalDecodeError, err)
	}

	// Process
	fmt.Println("GetData:2")

	var idAttrib dynamodb.AttributeValue
	idAttrib.S = &rqt.Id

	var get dynamodb.GetItemInput
	get.TableName = &tableName
	get.Key = make(map[string]*dynamodb.AttributeValue)
	get.Key[tableAttID] = &idAttrib
	get.AttributesToGet = []*string{&tableAttName, &tableAttCoins}

	output, err := getDynamoConnection().GetItem(&get)
	if err != nil {
		return ase.AWSError(ase.Transmission_InternalHandlerError, err)
	}

	if output.Item == nil {
		return ase.AWSError(ase.Transmission_InternalHandlerError, errors.New("user doesn't exist"))
	}

	coins, err := strconv.Atoi(*output.Item[tableAttCoins].N)
	if err != nil {
		return ase.AWSError(ase.Transmission_InternalHandlerError, err)
	}

	resp := GetUserResponse{}
	resp.User.Id = rqt.Id
	resp.User.Name = *output.Item[tableAttName].S
	resp.User.Coins = int32(coins)

	return ase.AWSResponse(&resp)
}

// SetData is a function
func SetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("SetData:1")

	// Cross-Origin Domain - WebGL
	respOptions, shouldStop := ase.AWSPreFlight(request)
	if shouldStop {
		return respOptions, nil
	}

	var rqt = SetUserRequest{}
	err := ase.AWSDecode(request, &rqt)
	if err != nil {
		return ase.AWSError(ase.Transmission_MarshalDecodeError, err)
	}

	fmt.Println("SetData:2")

	var idAttrib dynamodb.AttributeValue
	idAttrib.S = &rqt.User.Id

	var nameAttrib dynamodb.AttributeValue
	nameAttrib.S = &rqt.User.Name

	var coins = strconv.Itoa(int(rqt.User.Coins))
	var coinsAttrib dynamodb.AttributeValue
	coinsAttrib.N = &coins

	var put dynamodb.PutItemInput
	put.TableName = &tableName
	put.Item = make(map[string]*dynamodb.AttributeValue)
	put.Item[tableAttID] = &idAttrib
	put.Item[tableAttName] = &nameAttrib
	put.Item[tableAttCoins] = &coinsAttrib

	_, err = getDynamoConnection().PutItem(&put)
	if err != nil {
		return ase.AWSError(ase.Transmission_InternalHandlerError, err)
	}

	resp := SetUserResponse{}
	resp.HasCreated = true

	return ase.AWSResponse(&resp)
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
