package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"interest-arbitrage/constant"
	"interest-arbitrage/global"
	"interest-arbitrage/model"
	"interest-arbitrage/model/entity"
	"interest-arbitrage/router"
	"interest-arbitrage/server"
	"interest-arbitrage/utils"
	"log"
	"strconv"
	"time"
)

func main() {
	initConfig()
	initEnv()
	initDB()
	_ = initNode()
	// initPrice(node)
	initABI()
	initRouter()
}

func initABI() {
	global.ContractInfoMap = make(map[string]*model.ContractInfo)
	for _, info := range constant.ContractInfos {
		contractABIByFile, err := utils.GetContractABIByFile(info.ABIPath)
		if err != nil {
			log.Fatalf("fialed to get abi, path: %s, err: %s", info.ABIPath, err)
		}
		info.ABI = contractABIByFile
		global.ContractInfoMap[info.Name] = info
	}
}

func initRouter() {
	r := gin.Default()
	// 注册路由
	router.RegisterRouter(r)
	var addr string
	if global.Config.Server.Port == 0 {
		addr = ":8080"
	} else {
		addr = ":" + strconv.Itoa(int(global.Config.Server.Port))
	}
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}

func initPrice(node *server.NodeServer) {
	priceServer := server.GetPriceServer(node)
	priceServer.Start()
	server.UsablePriceServer = priceServer
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
	nodeServer := server.GetDefaultNodeServer(global.Env.NodeUrl)
	nodeServer.Start()
	server.UsableNodeServer = nodeServer
	return nodeServer
}

func initDB() {
	// gorm日志
	// gormLogger := logger.New(
	// 	log.New(os.Stdout, "sql: ", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // 慢SQL阈值
	// 		LogLevel:      logger.Info,
	// 		Colorful:      true,
	// 	},
	// )
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键
		// Logger:                                   gormLogger,
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
	syncTable()
}

func syncTable() {
	err := global.DB.AutoMigrate(
		&entity.ChainLinkConfig{},
		&entity.TokenPrice{},
		&entity.UniswapV2PriceConfig{},
		&entity.UniswapV2TokenCumulativeLast{},
		&entity.Price{},
		&entity.ChainInfo{},
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
