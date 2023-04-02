package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"new_shop_srv/goods_srv/global"
	"os"
	"time"
)

func SqlClientInit() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	dsn := fmt.Sprintf("root:root@tcp(%v:%v)/shop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local", global.NacosConf.MysqlConfig.Host, global.NacosConf.MysqlConfig.Port)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			//LogLevel:      logger.Silent, // Log level
			LogLevel: logger.Info,
			Colorful: false, // 禁用彩色打印
		},
	)
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "",
			SingularTable: true,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		panic(err)
	}
	//err = DB.AutoMigrate(&model.Brands{})
	//err = DB.AutoMigrate(&model.Category{})
	//err = DB.AutoMigrate(&model.Goods{})
	//err = DB.AutoMigrate(&model.Banner{})
	//err = DB.AutoMigrate(&model.GoodsCategoryBrand{})
	if err != nil {
		panic(err)
	}
	// 赋值给全局变量
	global.DB = DB
	// 制造数据
	//options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//dbPassWord := fmt.Sprintf("$pbkdf2-sha256$%s$%s", salt, encodedPwd)
	//
	//for i := 0; i <= 10; i++ {
	//	user := model.User{}
	//	user.NickName = fmt.Sprintf("zcc%d", i)
	//	user.Mobile = fmt.Sprintf("1923412231%d", i)
	//	user.Password = dbPassWord
	//	DB.Save(&user)
	//}
}
