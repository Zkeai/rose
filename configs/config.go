package configs

import (
	"errors"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type App struct {
	PrefixUrl     string
	ServerAddress string
	AppName       string

	LogSavePath string
	LogSaveName string
	LogFileExt  string

	MaxLogFiles int

	ImageStaticPath string
	ImageSavePath   string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Mysql struct {
	Host     string
	Port     string
	Dbname   string
	User     string
	Password string
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

type Jwt struct {
	Secret         string
	ExpirationDays int
}

type Config struct {
	App    App
	Server Server
	Mysql  Mysql
	Jwt    Jwt
	Redis  Redis
}

var C Config

// Setup initialize the configuration instance
func Setup() {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatalf("Config file not found")
		}
	}

	viper.AutomaticEnv()

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	applyEnvVariables()
}

func applyEnvVariables() {
	C.App.PrefixUrl = viper.GetString("PREFIX_URL")
	C.App.ServerAddress = viper.GetString("SERVER_ADDRESS")
	C.App.AppName = viper.GetString("APP_NAME")
	C.Server.RunMode = viper.GetString("RUN_MODE")
	C.Server.HttpPort = viper.GetInt("HTTP_PORT")
	C.Mysql.Host = viper.GetString("MYSQL_HOST")
	C.Mysql.Port = viper.GetString("MYSQL_PORT")
	C.Mysql.Password = viper.GetString("MYSQL_PASSWORD")
	C.Mysql.Dbname = viper.GetString("MYSQL_DBNAME")
	C.Mysql.User = viper.GetString("MYSQL_USER")
	C.Redis.Addr = viper.GetString("REDIS_ADDR")
	C.Redis.Password = viper.GetString("REDIS_PASSWORD")
	C.Redis.DB = viper.GetInt("REDIS_DB")
	C.Jwt.Secret = viper.GetString("JWT_SECRET")
	C.Jwt.ExpirationDays = viper.GetInt("JWT_EXPIRATION_DAYS")
}
