package userstorage

import (
	"context"
	"greport/common"
	usermodel "greport/module/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*usermodel.Users, error) {
	var returnData usermodel.Users
	if err := s.db.Where(condition).First(&returnData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &returnData, nil
}
