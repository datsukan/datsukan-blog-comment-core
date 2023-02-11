package model

// Comment は、コメントを保持するテーブルの構造体。
type Comment struct {
	ID        string `dynamo:"ID,hash"`
	ArticleID string `dynamo:"ArticleID"`
	ParentID  string `dynamo:"ParentID"`
	UserName  string `dynamo:"UserName"`
	Content   string `dynamo:"Content"`
	CreatedAt string `dynamo:"CreatedAt"`
}
