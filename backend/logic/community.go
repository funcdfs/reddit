package logic

import (
	"reddit/dao/mysql"
	"reddit/models"
)

// GetCommunityList use mysql.GetCommunityList to return a list of communityList from the database
func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库中所有的 community 并返回
	return mysql.GetCommunityList()
}
