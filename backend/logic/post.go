package logic

import (
	"reddit/dao/mysql"

	"go.uber.org/zap"

	"reddit/models"
	"reddit/pkg/gen"
)

func CreatePost(post *models.Post) (err error) {
	// 1、 生成 post_id(生成帖子ID)
	postID := gen.NewID()
	post.PostID = uint64(postID)

	// 2、创建帖子 保存到数据库
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	// community, err := mysql.GetCommunityNameByID(fmt.Sprint(post.CommunityID))
	// if err != nil {
	// 	zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
	// 	return err
	// }
	//
	// // redis 存储帖子信息
	// if err := redis.CreatePost(
	// 	post.PostID,
	// 	post.AuthorId,
	// 	post.Title,
	// 	TruncateByWords(post.Content, 120),
	// 	community.CommunityID); err != nil {
	// 	zap.L().Error("redis.CreatePost failed", zap.Error(err))
	// 	return err
	// }
	return

}

// GetPostById search post detail by post id
func GetPostById(id int64) (data *models.Post, err error) {
	return mysql.GetPostByID(id)
}
