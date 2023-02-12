package repoif

import "github.com/datsukan/datsukan-blog-comment-core/model"

// BlogGoodRepository は、BlogGood テーブルを操作するための Repository インターフェイス。
type CommentRepository interface {
	ReadByArticleID(articleID string) ([]*model.Comment, error)
	Create(articleID string, parentID string, name string, content string) (*model.Comment, error)
}
