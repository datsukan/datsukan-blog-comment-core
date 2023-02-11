package pkg_test

import (
	"testing"

	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/stretchr/testify/assert"
)

func Test_NewDynamoDBClient(t *testing.T) {
	t.Run("DynamoDBのクライアントインスタンスが正常に生成できること", func(t *testing.T) {
		assert := assert.New(t)

		db, err := pkg.NewDynamoDBClient()
		assert.NoError(err)
		assert.NotNil(db)
	})
}
