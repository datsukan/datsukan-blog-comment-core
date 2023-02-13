package usecase

import (
	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/datsukan/datsukan-blog-comment-core/repo"
	"github.com/datsukan/datsukan-blog-comment-core/repoif"
)

// Count は、指定された記事に紐づくコメントの件数を取得する。
func Count(articleID string) (int64, error) {
	db, err := pkg.NewDynamoDBClient()
	if err != nil {
		return 0, err
	}

	r := repoif.CommentRepository(repo.NewCommentRepository(db))

	c, err := r.CountByArticleID(articleID)
	if err != nil {
		return 0, err
	}

	return c, nil
}
