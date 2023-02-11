package pkg_test

import (
	"testing"

	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	t.Run("SQSのクライアントインスタンスが正常に生成できること", func(t *testing.T) {
		assert := assert.New(t)

		sqs := pkg.NewSQSClient()
		assert.NotNil(sqs)
	})
}
