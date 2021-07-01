package mysql

import (
	"bluebell_demo/models"
	"database/sql"
	"time"

	"go.uber.org/zap"
)

type comment1 struct {
	Post_id     uint64    `db:"post_id" json:"post_id"`
	Parent_id   uint64    `db:"parent_id" json:"parent_id"`
	Comment_id  uint64    `db:"comment_id" json:"comment_id"`
	Author_id   uint64    `db:"author_id" json:"author_id"`
	Content    string    `db:"content" json:"content"`
	Create_time time.Time `db:"create_time" json:"create_time"`
	Username   string
}
func CreateComment(comment *models.Comment) (err error) {
	sqlStr := `insert into comment(comment_id, content, post_id, author_id, parent_id) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID,
		comment.AuthorID, comment.ParentID)
	if err != nil {
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetCommentListByIDs(ids int64) (commentList []comment1, err error) {
	var comments []comment1
	sqlStr := `select comment_id, content, post_id, author_id, parent_id, create_time from comment where post_id = ?`
	err = db.Select(&comments,sqlStr,ids)
	if err == sql.ErrNoRows {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	commentList=comments
	return
}
