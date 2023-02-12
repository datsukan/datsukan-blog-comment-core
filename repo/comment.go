package repo

import (
	"fmt"
	"time"

	"github.com/datsukan/datsukan-blog-comment-core/model"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

const timeformat = "2006-01-02 15:04:05.000000000"

// CommentRepository は、DynamoDB 用の DB の構造体。
type CommentRepository struct {
	Table dynamo.Table
}

// NewCommentRepository は、 CommentRepository のインスタンスを生成する。
func NewCommentRepository(db *dynamo.DB) *CommentRepository {
	return &CommentRepository{Table: db.Table("DatsukanBlogComment")}
}

// ReadByArticleID は、記事に紐づくコメントの一覧を取得する。
func (r *CommentRepository) ReadByArticleID(articleID string) ([]*model.Comment, error) {
	var cs []*model.Comment
	err := r.Table.Get("ArticleID", articleID).Order(dynamo.Ascending).All(&cs)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
		return nil, err
	}

	return cs, nil
}

// Create は、コメントのレコードを作成する。
func (r *CommentRepository) Create(articleID string, parentID string, userName string, content string) (*model.Comment, error) {
	c := &model.Comment{
		ID:        uuid.New().String(),
		ArticleID: articleID,
		ParentID:  parentID,
		UserName:  userName,
		Content:   content,
		CreatedAt: time.Now().Format(timeformat),
	}

	if err := r.Table.Put(c).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
		return nil, err
	}

	return c, nil
}
