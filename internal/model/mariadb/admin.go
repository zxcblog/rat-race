package mariadb

type Admin struct {
	Model
	Username    string `gorm:"column:username;not null;comment:用户名" json:"username"`
	Password    string `gorm:"column:password;not null;comment:密码" json:"password"`
	Nickname    string `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`
	Mobile      string `gorm:"column:mobile;not null;comment:手机号" json:"mobile"`
	CreatedUser string `gorm:"column:created_user;not null;comment:创建人" json:"created_user"`
	UpdatedUser string `gorm:"column:updated_user;not null;comment:更新人" json:"updated_user"`
	Status      uint   `gorm:"type:tinyint(1);column:status;not null;default:1;comment:状态 1:正常  2:禁用" json:"status"`
}
