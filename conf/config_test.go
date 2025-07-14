package conf

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	// 执行初始化
	InitConfig()

	// 验证配置是否正确加载
	if Config == nil {
		t.Fatal("Config is nil after initialization")
	}

	// 验证System配置
	if Config.System == nil {
		t.Fatal("System config is nil")
	}
	if Config.System.Domain != "mall" {
		t.Errorf("Expected System.Domain to be 'mall', got '%s'", Config.System.Domain)
	}
	if Config.System.AppEnv != "dev" {
		t.Errorf("Expected System.AppEnv to be 'dev', got '%s'", Config.System.AppEnv)
	}
	if Config.System.HttpPort != "5001" {
		t.Errorf("Expected System.HttpPort to be '5001', got '%s'", Config.System.HttpPort)
	}

	// 验证Mysql配置
	if Config.Mysql == nil {
		t.Fatal("Mysql config is nil")
	}
	defaultMysql, ok := Config.Mysql["default"]
	if !ok {
		t.Fatal("Mysql default config not found")
	}
	if defaultMysql.DbHost != "rm-3nssj3q44i71c190q.mysql.rds.aliyuncs.com" {
		t.Errorf("Expected Mysql default dbHost to be 'rm-3nssj3q44i71c190q.mysql.rds.aliyuncs.com', got '%s'", defaultMysql.DbHost)
	}
	if defaultMysql.UserName != "user_rw" {
		t.Errorf("Expected Mysql default userName to be 'user_rw', got '%s'", defaultMysql.UserName)
	}

	// 验证Redis配置
	if Config.Redis == nil {
		t.Fatal("Redis config is nil")
	}
	if Config.Redis.RedisHost != "r-3ns7huz6q0teloxquw.redis.rds.aliyuncs.com" {
		t.Errorf("Expected Redis redisHost to be 'r-3ns7huz6q0teloxquw.redis.rds.aliyuncs.com', got '%s'", Config.Redis.RedisHost)
	}
	if Config.Redis.RedisPassword != "front_risk_rw:E4gb2ji580" {
		t.Errorf("Expected Redis redisPassword to be 'front_risk_rw:E4gb2ji580', got '%s'", Config.Redis.RedisPassword)
	}
}
