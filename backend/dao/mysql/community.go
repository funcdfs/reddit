package mysql

import (
	"database/sql"

	"go.uber.org/zap"

	"reddit/models"
)

// GetCommunityList is the basic function to get the communityList from the database
// the error is nil, just log in the console
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community list in database")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID
// use a id to get specify community detail
func GetCommunityDetailByID(id uint64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInValidID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorInValidID
	}
	return community, err
}

func GetCommunityByID(id uint64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows { // 查询为空
		err = ErrorInValidID // 无效的ID
		return
	}
	if err != nil {
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryError
	}
	return community, err
}
