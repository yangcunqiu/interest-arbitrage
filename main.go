package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"interest-arbitrage/global"
	"interest-arbitrage/model"
	"interest-arbitrage/server"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	initConfig()
	initEnv()
	initDB()
	node := initNode()
	initPrice(node)
	time.Sleep(time.Hour)
}

func initPrice(node *server.NodeServer) {
	priceServer := server.GetPriceServer(node)
	priceServer.Start()
}

func initEnv() {
	// 设置配置文件名称 类型
	viper.SetConfigName("env")
	viper.SetConfigType("yml")
	// 设置viper的查找路径(viper会去这个路径查找上面配置的文件名的文件)
	viper.AddConfigPath(".")
	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 把配置信息反序列化到结构体, 方便使用
	err = viper.Unmarshal(&global.Env)
	if err != nil {
		log.Println(err)
	}
	// 运行时监控配置文件的更新
	viper.WatchConfig()
	// 配置文件更新时的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config changed:", e.Name)
		// 重新反序列化
		err = viper.Unmarshal(&global.Env)
		if err != nil {
			log.Println(err)
		}
	})
}

func initNode() *server.NodeServer {
	nodeServer := server.GetDefaultNodeServer(global.Env.AlchemyUrl)
	nodeServer.Start()
	return nodeServer
}

func initDB() {
	// gorm日志
	gormLogger := logger.New(
		log.New(os.Stdout, "/r/n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键
		Logger:                                   gormLogger,
	})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConnCount)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConnCount)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifetime) * time.Millisecond)
	global.DB = db
	// 同步表结构
	//syncTable()
}

func syncTable() {
	err := global.DB.AutoMigrate(
		&model.ChainLinkConfig{},
		&model.TokenPrice{},
	)
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	// 设置配置文件名称 类型
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	// 设置viper的查找路径(viper会去这个路径查找上面配置的文件名的文件)
	viper.AddConfigPath(".")
	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 把配置信息反序列化到结构体, 方便使用
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		log.Println(err)
	}
	// 运行时监控配置文件的更新
	viper.WatchConfig()
	// 配置文件更新时的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config changed:", e.Name)
		// 重新反序列化
		err = viper.Unmarshal(&global.Config)
		if err != nil {
			log.Println(err)
		}
	})
}

func getDsn() string {
	// dsn := "root:pass1234@tcp(106.14.18.18:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	m := global.Config.Mysql
	if m.Username == "" {
		panic("数据库用户名不能为空")
	}
	if m.Password == "" {
		panic("数据库密码不能为空")
	}
	if m.Ip == "" {
		panic("数据库IP不能为空")
	}
	if m.Database == "" {
		panic("数据库库名不能为空")
	}
	if m.Port == 0 {
		m.Port = 3306
	}
	var params string
	if m.Params != "" {
		params = "?" + m.Params
	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Ip + ":" + strconv.Itoa(int(m.Port)) + ")" + "/" + m.Database + params
}
