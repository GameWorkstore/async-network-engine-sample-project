package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	Name string `json:"name"`
}

// GetData is a function
func GetData(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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

func main() {

	var functionName = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	fmt.Println("FunctionName" + functionName)
	switch functionName {
	case "writedata":
		lambda.Start(ImplementationFailure)
		return
	case "lksgetdata":
		lambda.Start(GetData)
		return
	default:
		lambda.Start(ImplementationFailure)
		return
	}
}
