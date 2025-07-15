package conf

import (
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	System *System `yaml:"system"`

	MySql         map[string]*Mysql `yaml:"mysql"`
	Redis         *Redis            `yaml:"redis"`
	EncryptSecret *EncryptSecret    `yaml:"encryptSecret"`
}

type System struct {
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	AppEnv      string `yaml:"env"`
	HttpPort    string `yaml:"HttpPort"`
	Host        string `yaml:"Host"`
	UploadModel string `yaml:"UploadModel"`
}

type Mysql struct {
	Dialect         string `yaml:"dialect"`
	DbHost          string `yaml:"dbHost"`
	DbPort          string `yaml:"dbPort"`
	DbName          string `yaml:"dbName"`
	UserName        string `yaml:"userName"`
	Password        string `yaml:"password"`
	Charset         string `yaml:"charset"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type Redis struct {
	RedisDbName   string `yaml:"redisDbName"`
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisPassword string `yaml:"redisPassword"`
}

// EncryptSecret 加密的东西
type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
	MoneySecret string `yaml:"moneySecret"`
}

// InitConfig 读取当前环境配置文件
func InitConfig() {
	workDir, _ := os.Getwd()      // 获取当前工作目录
	viper.SetConfigName("config") // 配置文件名（不含扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型

	viper.AddConfigPath(workDir + "/conf") // 配置文件搜索路径1
	viper.AddConfigPath(workDir)           // 配置文件搜索路径2

	err := viper.ReadInConfig() //读取配置文件
	if err != nil {
		panic(err) // 如果读取失败，直接panic
	}

	err = viper.Unmarshal(&Config) // 将配置解析搭配全局变量Config
	if err != nil {
		panic(err)
	}
}
