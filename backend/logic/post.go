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
func GetPostById(id int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	// 查询帖子信息
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed",
			zap.Int64("postID", id),
			zap.Error(err))
		return nil, err
	}

	// 根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed",
			zap.Uint64("postID", post.AuthorId),
			zap.Error(err))
		return
	}

	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed",
			zap.Uint64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	// 接口数据拼接
	data = &models.ApiPostDetail{
		Post:            post,
		CommunityDetail: community,
		AuthorName:      user.UserName,
	}
	return
}
