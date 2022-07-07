package models

import "time"

type Post struct {
	PostID      uint64    `json:"post_id,string" db:"post_id"`
	AuthorId    uint64    `json:"author_id" db:"author_id"`
	CommunityID uint64    `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

// 尽量把相同结构的类型塞到一起

// ApiPostDetail define the specific parameters in the ApiPostDetail structures
type ApiPostDetail struct {
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
	// CommunityName string `json:"community_name"`
}
