package userstorage

import (
	"2ndbrand-api/common"
	usermodel "2ndbrand-api/module/user/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, dataModel *usermodel.UserCreate) error {
	db := s.db.Begin()
	if err := db.Table(dataModel.TableName()).Create(&dataModel).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
