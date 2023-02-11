package response

import (
	"github.com/aws/aws-lambda-go/events"
)

// ResponseSuccess は処理成功時のレスポンスを生成する
func ResponseSuccess(body string) (events.APIGatewayProxyResponse, error) {
	r := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
	}
	return r, nil
}
