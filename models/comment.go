package models

import "C"
import (
	"encoding/json"
	"errors"
	"time"
)

type Comment struct {
	PostID     uint64    `db:"question_id" json:"question_id,string"`
	ParentID   uint64    `db:"parent_id" json:"parent_id,string"`
	CommentID  uint64    `db:"comment_id" json:"comment_id,string"`
	AuthorID   uint64    `db:"author_id" json:"author_id,string"`
	Content    string    `db:"content" json:"content"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}


func (c *Comment) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID uint64 `db:"question_id" json:"question_id,string"`
		ParentID uint64 `db:"parent_id" json:"parent_id,string"`
		Content    string    `db:"content" json:"content"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if required.PostID == 0 {
		err = errors.New("缺少必填字段PostID")
	} else if required.ParentID == 0 {
		err = errors.New("缺少必填字段ParentID")
	}else if len(required.Content)==0 {
		err = errors.New("缺少必填字段Content")
	} else {
		c.PostID = required.PostID
		c.ParentID = required.ParentID
		c.Content=required.Content
	}
	return
}