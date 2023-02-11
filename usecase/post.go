package usecase

import (
	"fmt"

	"github.com/datsukan/datsukan-blog-comment-core/model"
	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/datsukan/datsukan-blog-comment-core/repo"
	"github.com/datsukan/datsukan-blog-comment-core/repoif"
)

func Post(articleID string, parentID string, userName string, content string) (*model.Comment, error) {
	db, err := pkg.NewDynamoDBClient()
	if err != nil {
		return nil, err
	}

	r := repoif.CommentRepository(repo.NewCommentRepository(db))

	c, err := r.Create(articleID, parentID, userName, content)
	if err != nil {
		return nil, err
	}

	fmt.Printf("ID: %s, ArticleID: %s, ParentID: %s, UserName: %s, Content: %s, CreatedAt: %s\n", c.ID, c.ArticleID, c.ParentID, c.UserName, c.Content, c.CreatedAt)
	return c, nil
}
