package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/datsukan/datsukan-blog-comment-core/model"
	"github.com/datsukan/datsukan-blog-comment-core/pkg"
	"github.com/datsukan/datsukan-blog-comment-core/repo"
	"github.com/datsukan/datsukan-blog-comment-core/repoif"
)

// Post は、指定された内容でコメントを登録する。
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

	input := NoticeInput{
		ArticleID: c.ArticleID,
		CommentID: c.ID,
		UserName:  c.UserName,
		Comment:   c.Content,
	}
	noticeEnqueue(input)

	fmt.Printf("ID: %s, ArticleID: %s, ParentID: %s, UserName: %s, Content: %s, CreatedAt: %s\n", c.ID, c.ArticleID, c.ParentID, c.UserName, c.Content, c.CreatedAt)
	return c, nil
}

type NoticeInput struct {
	ArticleID string `json:"article_id"`
	CommentID string `json:"comment_id"`
	UserName  string `json:"user_name"`
	Comment   string `json:"comment"`
}

// noticeEnqueue は、通知用のSQSにメッセージを追加する。
func noticeEnqueue(input NoticeInput) {
	queueURL := os.Getenv("QUEUE_URL")
	sqsSvc := pkg.NewSQSClient()

	msgJson, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = sqsSvc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(msgJson)),
		QueueUrl:    &queueURL,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("notification successful")
}
