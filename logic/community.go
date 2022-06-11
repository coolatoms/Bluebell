package logic

import (
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/models"
)

// GetCommunityList 查询逻辑
func GetCommunityList() ([]*models.Community, error) {
	//	查community数据库返回相应值
	return mysql.GetCommunityList()
}

// GetCommunityDetailList 获取communityID逻辑
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
