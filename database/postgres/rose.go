package postgres

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// EnvRoseDb ..
type EnvRoseDb struct {
	User           string `envconfig:"ROSE_POSTGRES_USER"`
	Pass           string `envconfig:"ROSE_POSTGRES_PASS"`
	Name           string `envconfig:"ROSE_POSTGRES_NAME"`
	Host           string `envconfig:"ROSE_POSTGRES_HOST"`
	Port           string `envconfig:"ROSE_POSTGRES_PORT"`
	Debug          bool   `envconfig:"ROSE_POSTGRES_DEBUG" default:"true"`
	Type           string `envconfig:"ROSE_TYPE" default:"postgres"`
	SslMode        string `envconfig:"ROSE_POSTGRES_SSL_MODE" default:"disable"`
	ConnectTimeout string `envconfig:"ROSE_POSTGRES_TIMEOUT"`
}

var (
	// DbConRose ..
	DbConRose *gorm.DB
	// DbRoseErr ..
	DbRoseErr error
	envRoseDb EnvRoseDb
)

func init() {

	fmt.Println("DB ROSE")

	err := envconfig.Process("ROSE", &envRoseDb)
	if err != nil {
		fmt.Println("Failed to get Otto SFA Callplan env:", err)
	}

	dbopen, err := DbRoseOpen()
	if err != nil {
		fmt.Println("Can't open", envRoseDb.Name, "DB")
	}

	DbConRose = dbopen

}

// DbRoseOpen ..
func DbRoseOpen() (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", envRoseDb.Host, envRoseDb.Port, envRoseDb.User, envRoseDb.Pass, envRoseDb.Name, envRoseDb.SslMode, envRoseDb.ConnectTimeout)
	// DbCon, DbErr = gorm.Open("postgres", args)
	DbConRose, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		logs.Error("Open", envRoseDb.Name, "DB error :", err)
		return DbConRose, err
	}

	sqlDB, err := DbConRose.DB()
	if err != nil {
		logs.Error("SQL DB not connected :", err)
		return DbConRose, err
	}
	if errping := sqlDB.Ping(); errping != nil {
		logs.Error("DB not connected test ping :", errping)
		return DbConRose, errping
	}

	return DbConRose, nil
}
