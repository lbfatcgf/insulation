package service

import (
	"insulation/server/admin/internal/repositorys/sys_user_rep"
	"insulation/server/base/pkg/database"
	hashutil "insulation/server/base/pkg/hash_util"
	"insulation/server/base/pkg/models"
	"insulation/server/base/pkg/password"
)

func InitSys() {
	database.Initialize()
	db := database.GetDB()
	db.AutoMigrate(models.SysUser{})
	count, err := sys_user_rep.CountAll()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		err = createSuperAdmin()
		if err != nil {
			panic(err)
		}
	}
}

func createSuperAdmin() error {
	psword, err := password.Gen(hashutil.Sm3String("admin123456"))
	if err != nil {
		return err
	}
	user := models.SysUser{
		UserName:    "admin",
		Email:       "insulation@admin.com",
		PhoneNumber: "8613322222222",
		Password:    psword,
	}
	err = sys_user_rep.Create(&user)

	return err
}
