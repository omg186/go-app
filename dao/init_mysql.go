package dao

import (
	"fmt"
	"order-food-service/entitys"
	"order-food-service/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 定义自己的Writer
type CustomWriter struct {
	customLogger *logrus.Logger
}

//实现gorm/logger.Writer接口
func (c *CustomWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	c.customLogger.Info(logstr)
}
func NewCustomWriter() *CustomWriter {
	log := global.AppLog
	//配置logrus
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return &CustomWriter{customLogger: log}
}

func Init(MYSQL_USER_NAME string, MYSQL_PASSWORD string, MYSQL_URL string, MYSQL_PORT string, MYSQL_DATABASE string, MYSQL_MAX_IDLE_CONNS int, MYSQL_MAX_OPEN_CONNS int) *gorm.DB {
	//dsn := "root:yungengny@tcp(127.0.0.1:3306)/cms_farming?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL_USER_NAME + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_URL + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	global.AppLog.Info("dsn = ", dsn)

	slowLogger := logger.New(
		//设置Logger
		NewCustomWriter(),
		logger.Config{
			//慢SQL阈值
			SlowThreshold: time.Millisecond,
			//设置日志级别，只有Info以上才会打印sql
			LogLevel: logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   slowLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止外键约束
	})
	if err != nil {
		global.AppLog.Error(err)
	}

	sqlDb, _ := db.DB()
	// db.Re
	sqlDb.SetMaxIdleConns(MYSQL_MAX_IDLE_CONNS) //设置最大连接数
	sqlDb.SetMaxOpenConns(MYSQL_MAX_OPEN_CONNS) //设置最大的空闲连接数
	db.AutoMigrate(&entitys.User{})
	return db
}
