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
