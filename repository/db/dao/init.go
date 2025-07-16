package dao

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go-mall/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// InitMySQL 初始化数据库连接池
func InitMySQL() error {
	var initErr error
	dbOnce.Do(func() {
		mConfig := conf.Config.MySql["default"]

		// 构建DSN
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
			mConfig.UserName,
			mConfig.Password,
			mConfig.DbHost,
			mConfig.DbPort,
			mConfig.DbName,
			mConfig.Charset)

		// 初始化GORM配置
		gormConfig := &gorm.Config{
			Logger: getLogger(),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			PrepareStmt:            true, // 开启预编译语句缓存
			SkipDefaultTransaction: true, // 禁用默认事务
		}

		// 创建数据库连接
		var err error
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), gormConfig)

		if err != nil {
			initErr = fmt.Errorf("failed to connect to database: %v", err)
			return
		}

		// 配置连接池
		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("failed to get sql.DB: %v", err)
			return
		}

		sqlDB.SetMaxIdleConns(mConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mConfig.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(mConfig.ConnMaxLifetime) * time.Second)

		// 设置表选项
		db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")

		// 数据库自动建表以及迁移
		if err := migrate(); err != nil {
			initErr = fmt.Errorf("failed to migrate database: %v", err)
			return
		}

		// 启动健康检查
		go checkDBHealth(db, mConfig)
	})

	return initErr
}

// getLogger 获取GORM日志配置
func getLogger() logger.Interface {
	if gin.Mode() == gin.DebugMode {
		return logger.Default.LogMode(logger.Info)
	}
	return logger.Default.LogMode(logger.Warn)
}

// checkDBHealth 定期检查数据库健康状态
func checkDBHealth(db *gorm.DB, config *conf.Mysql) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Failed to get sql.DB: %v", err)
			continue
		}

		if err := sqlDB.Ping(); err != nil {
			log.Printf("Database connection is down: %v", err)
			// 这里可以添加重连逻辑
		}
	}
}

// NewDBClient 获取数据库客户端
func NewDBClient(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
