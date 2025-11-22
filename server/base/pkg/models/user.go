package models

type SysUser struct {
	BaseModel
	UserName    string `gorm:"column:name;type:text;not null" json:"userName"`
	Email       string `gorm:"column:email;type:text;not null" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;type:text" json:"phoneNumber"`
	Password    string `gorm:"column:password;type:text;not null" json:"password"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
