package models

import (
	"github.com/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// gorm 进行mysql连接 返回DB+error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{ //建立数据库链接
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}
	//自动根据结构体创建表 判断是否返回error
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{}) //初始化表  用户信息 视频 评论 用户登录
	if err != nil {
		panic(err)
	}
}
