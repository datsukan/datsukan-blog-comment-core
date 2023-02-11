package response_test

import (
	"fmt"
	"testing"

	"github.com/datsukan/datsukan-blog-comment-core/response"
	"github.com/stretchr/testify/assert"
)

func Test_ResponseBadRequestError(t *testing.T) {
	t.Run("戻り値が、404 バッドリクエスト用のAPIGatewayのレスポンス構造体であること。", func(t *testing.T) {
		assert := assert.New(t)

		r, err := response.ResponseBadRequestError(fmt.Errorf("test 404"))
		assert.NoError(err)
		assert.Equal(r.StatusCode, 400)
		assert.Equal(r.Body, "{\"error\":\"bad request\",\"message\":\"test 404\"}")
	})
}

func Test_ResponseInternalServerError(t *testing.T) {
	t.Run("戻り値が、500 サーバーエラー用のAPIGatewayのレスポンス構造体であること。", func(t *testing.T) {
		assert := assert.New(t)

		r, err := response.ResponseInternalServerError(fmt.Errorf("test 500"))
		assert.NoError(err)
		assert.Equal(r.StatusCode, 500)
		assert.Equal(r.Body, "{\"error\":\"internal server error\",\"message\":\"test 500\"}")
	})
}
