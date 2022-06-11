package mysql

import (
	"database/sql"
	"studyWeb/Bluebell/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlstr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlstr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("this is no db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 获取社区帖子ID
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlstr := `select community_id,community_name,introduction,create_time from community where community_id=?`
	err = db.Get(community, sqlstr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
