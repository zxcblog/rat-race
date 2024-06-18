package mariadb

type User struct {
	ModelID

	Username        string `gorm:"column:username;not null;comment:用户名" json:"username"`
	Password        string `gorm:"column:password;not null;comment:密码" json:"password"`
	Email           string `gorm:"column:email;not null;comment:邮箱" json:"email"`
	Mobile          string `gorm:"column:mobile;not null;comment:手机号" json:"mobile"`
	EmailVerifiedAt uint32 `gorm:"column:email_verified_at;not null;comment:邮箱验证时间" json:"email_verified_at"`
	Status          uint   `gorm:"type:tinyint(1);column:status;not null;default:1;comment:状态 1:正常  2:禁用" json:"status"`

	ModelTime
	ModelDelete
}
