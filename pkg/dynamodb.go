package pkg

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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
	conf := &aws.Config{
		Region: aws.String(dynamoDbRegion),
	}

	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	isLocal := len(dynamoDbEndpoint) != 0
	if isLocal {
		conf.Endpoint = aws.String(dynamoDbEndpoint)
		conf.DisableSSL = aws.Bool(true)
		conf.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "dummy")
	}

	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, err
	}

	return dynamo.New(sess), nil
}
