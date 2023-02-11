package pkg

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// NewDynamoDBClient は、 DynamoDB のクライアントを生成する。
func NewDynamoDBClient() (*dynamo.DB, error) {
	dynamoDbRegion := os.Getenv("AWS_REGION")
	// デフォルトでは東京リージョンを指定する。
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	// DynamoDB Local を利用する場合は Endpoint の URL を設定する。
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")

	disableSsl := false
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	sess, err := session.NewSession(&aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})
	if err != nil {
		return nil, err
	}

	return dynamo.New(sess), nil
}
