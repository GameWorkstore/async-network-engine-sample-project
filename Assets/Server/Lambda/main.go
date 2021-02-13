package appfunction

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// ImplementationFailure handles requests from API Gateway when function wasn't implemented properly
func ImplementationFailure(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var resp = events.APIGatewayProxyResponse{}
	resp.Body = "Implementation failure, verify if main contains the name of function"
	resp.StatusCode = 404

	return resp, nil
}

func main() {

	var functionName = os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
	switch functionName {
	case "writedata":
		lambda.Start(ImplementationFailure)
		return
	case "getdata":
		lambda.Start(ImplementationFailure)
		return
	default:
		lambda.Start(ImplementationFailure)
		return
	}
}
