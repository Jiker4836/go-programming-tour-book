package model

import (
	"fmt"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	// CreatedOn 创建时间
	CreatedOn uint32 `json:"created_on"`
	// CreatedBy 创建人
	CreatedBy string `json:"created_by"`
	// ModifiedOn 修改时间
	ModifiedOn uint32 `json:"modified_on"`
	// ModifiedBy 修改人
	ModifiedBy string `json:"modified_by"`
	// DeleteOn 删除时间
	DeleteOn uint32 `json:"delete_on"`
	// IsDel 是否删除 0-未删除 1-已删除
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?Charset=%s&parseTime=%t&loc=Local",
			databaseSetting.UserName,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
