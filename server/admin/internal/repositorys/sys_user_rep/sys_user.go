package sys_user_rep

import (
	"insulation/server/base/pkg/database"
	"insulation/server/base/pkg/logger"
	"insulation/server/base/pkg/models"
)

var log *logger.Logger

func CountAll() (int64, error) {
	if log == nil {
		log = logger.NewLogger("sys_user_repository", true)
	}
	var count int64
	res := database.GetDB().Table(models.SysUser{}.TableName()).Count(&count)
	if res.Error != nil {
		log.Error(res.Error.Error())
		return 0, res.Error
	}
	return count, nil
}

func QueryByLogin(phoneNumberOrEmail, password string) (*models.SysUser, error) {
	if log == nil {
		log = logger.NewLogger("sys_user_repository", true)
	}
	var user models.SysUser
	res := database.GetDB().First(
		&user,
		"(email = ? OR phone_number = ?) AND password = ?",
		phoneNumberOrEmail,
		phoneNumberOrEmail,
		password,
	)
	if res.Error != nil {
		log.Error(res.Error.Error())
		return nil, res.Error
	}
	return &user, nil
}

func Create(user *models.SysUser) error {
	if log == nil {
		log = logger.NewLogger("sys_user_repository", true)
	}

	res := database.GetDB().Create(user)
	if res.Error != nil {
		log.Error(res.Error.Error())
		return res.Error
	}
	return nil
}
