package response_test

import (
	"encoding/json"
	"testing"

	"github.com/datsukan/datsukan-blog-comment-core/response"
	"github.com/stretchr/testify/assert"
)

func Test_ResponseSuccess(t *testing.T) {
	t.Run("戻り値が、200 成功用のAPIGatewayのレスポンス構造体であること。", func(t *testing.T) {
		assert := assert.New(t)

		type Response struct {
			Hoge string `json:"hoge"`
			Fugo string `json:"fugo"`
		}
		b := Response{
			Hoge: "this is test",
			Fugo: "hello world",
		}
		jb, err := json.Marshal(b)
		assert.NoError(err)
		js := string(jb)

		r, err := response.ResponseSuccess(js)
		assert.NoError(err)
		assert.Equal(r.StatusCode, 200)
		assert.Equal(r.Body, "{\"hoge\":\"this is test\",\"fugo\":\"hello world\"}")
	})
}
