package usecase

import (
	"github.com/datsukan/datsukan-blog-comment-core/model"
	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/datsukan/datsukan-blog-comment-core/repo"
	"github.com/datsukan/datsukan-blog-comment-core/repoif"
)

type Comment struct {
	ID            string
	ArticleID     string
	ParentID      string
	UserName      string
	Content       string
	CreatedAt     string
	ReplyComments []*Comment
}

func Ref(articleID string) ([]*Comment, error) {
	db, err := pkg.NewDynamoDBClient()
	if err != nil {
		return nil, err
	}

	r := repoif.CommentRepository(repo.NewCommentRepository(db))

	comments, err := r.ReadByArticleID(articleID)
	if err != nil {
		return nil, err
	}

	var rcs []*Comment
	var ccs []*model.Comment
	for _, c := range comments {
		if c.ParentID != "" {
			ccs = append(ccs, c)
			continue
		}

		rc := &Comment{
			ID:        c.ID,
			ArticleID: c.ArticleID,
			ParentID:  c.ParentID,
			UserName:  c.UserName,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		}
		rcs = append(rcs, rc)
	}

	for _, c := range ccs {
		cc := &Comment{
			ID:        c.ID,
			ArticleID: c.ArticleID,
			ParentID:  c.ParentID,
			UserName:  c.UserName,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		}

		for _, rc := range rcs {
			if rc.ID == cc.ParentID {
				rc.ReplyComments = append(rc.ReplyComments, cc)
			}
		}
	}

	return rcs, nil
}
