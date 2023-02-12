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
	ReplyComments []*ReplyComment
}

type ReplyComment struct {
	ID        string
	ArticleID string
	ParentID  string
	UserName  string
	Content   string
	CreatedAt string
}

// Ref は、指定された記事に紐づくコメントのリストを取得する。
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

	rcs := []*Comment{}
	ccs := []*model.Comment{}
	for _, c := range comments {
		if c.ParentID != "" {
			ccs = append(ccs, c)
			continue
		}

		reply := []*ReplyComment{}
		rc := &Comment{
			ID:            c.ID,
			ArticleID:     c.ArticleID,
			ParentID:      c.ParentID,
			UserName:      c.UserName,
			Content:       c.Content,
			CreatedAt:     c.CreatedAt,
			ReplyComments: reply,
		}
		rcs = append(rcs, rc)
	}

	for _, c := range ccs {
		cc := &ReplyComment{
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
