package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response は異常系のレスポンスを定義した構造体
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ValidationErrorResponse は入力不正のレスポンスを定義した構造体
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ResponseBadRequestError はリクエスト不正のレスポンスを生成する
func ResponseBadRequestError(rerr error) (events.APIGatewayProxyResponse, error) {
	b := ErrorResponse{
		Error:   "bad request",
		Message: rerr.Error(),
	}
	jb, err := json.Marshal(b)
	if err != nil {
		r := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}
		return r, nil
	}
	body := string(jb)

	r := events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       body,
	}
	return r, nil
}

// ResponseInternalServerError はシステムエラーのレスポンスを生成する
func ResponseInternalServerError(rerr error) (events.APIGatewayProxyResponse, error) {
	b := ErrorResponse{
		Error:   "internal server error",
		Message: rerr.Error(),
	}
	jb, err := json.Marshal(b)
	if err != nil {
		r := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}
		return r, nil
	}
	body := string(jb)

	r := events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       body,
	}
	return r, nil
}
