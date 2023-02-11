package usecase

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-comment-core/model"
	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/datsukan/datsukan-blog-comment-core/repo"
	"github.com/datsukan/datsukan-blog-comment-core/repoif"
)

func Ref(articleID string) ([]*model.Comment, error) {
	db, err := pkg.NewDynamoDBClient()
	if err != nil {
		return nil, err
	}

	r := repoif.CommentRepository(repo.NewCommentRepository(db))

	comments, err := r.ReadByArticleID(articleID)
	if err != nil {
		return nil, err
	}

	for _, c := range comments {
		fmt.Printf("ID: %s, ArticleID: %s, ParentID: %s, UserName: %s, Content: %s, CreatedAt: %s\n", c.ID, c.ArticleID, c.ParentID, c.UserName, c.Content, c.CreatedAt)
	}
	return comments, nil
}
