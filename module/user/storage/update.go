package userstorage

import (
	"context"
	usermodel "greport/module/user/model"
)

func (store *sqlStore) Update(
	ctx context.Context,
	updateData *usermodel.UserUpdate,
	id int,
) error {
	if err := store.db.Where("id=?", id).Updates(&updateData); err != nil {
		return err.Error
	}
	return nil
}

func (store *sqlStore) UpdatePassword(
	ctx context.Context,
	updateData *usermodel.Users,
	id int,
) error {
	columnData := map[string]interface{}{
		"password": updateData.Password,
		"salt":     updateData.Salt,
	}
	if err := store.db.Table(updateData.TableName()).
		Where("id=?", id).UpdateColumns(&columnData); err != nil {
		return err.Error
	}
	return nil
}
