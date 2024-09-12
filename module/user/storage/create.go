package userstorage

import (
	"context"
	"greport/common"
	usermodel "greport/module/user/model"
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
