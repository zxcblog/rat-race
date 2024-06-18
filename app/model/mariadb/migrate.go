package mariadb

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Model struct {
	ModelID
	ModelTime
	ModelDelete
}

type ModelID struct {
	ID uint `gorm:"primary_key;column:id;comment:主键" json:"id"`
}

type ModelTime struct {
	CreatedAt uint32 `gorm:"autoCreateTime;column:created_at;comment:创建时间" json:"created_at"`
	UpdatedAt uint32 `gorm:"autoUpdateTime;column:updated_at;comment:更新时间" json:"updated_at"`
}

type ModelDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;comment:删除时间" json:"deleted_at"`
}

var migrate = []interface{}{
	&User{},
}

// InitSql mariadb数据库创建时初始化数据操作
func InitSql(db *gorm.DB) error {
	if len(migrate) < 1 {
		return nil
	}

	// 创建表信息
	if err := db.AutoMigrate(migrate...); err != nil {
		return errors.Wrap(err, "数据库表迁移失败")
	}

	//// 查看用户表信息是否存在， 不存在则添加
	//adminUser := Admin{Username: "race_admin"}
	//if err := db.First(&adminUser).Error; err != nil {
	//	if !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return err
	//	}
	//
	//	// 添加默认管理员
	//	adminUser = Admin{Username: "race_admin", Password: "race_admin", Nickname: "admin", Status: 1}
	//	if err = db.Create(&adminUser).Error; err != nil {
	//		return err
	//	}
	//}
	return nil
}
