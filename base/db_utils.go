package base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

type MysqlOptions struct {
	Username string
	Password string
	Ip       string
	DBName   string
}

type MysqlOption func(options *MysqlOptions)

func DBUsername(username string) MysqlOption {
	return func(options *MysqlOptions) {
		options.Username = username
	}
}

func DBPassword(password string) MysqlOption {
	return func(options *MysqlOptions) {
		options.Password = password
	}
}

func DBIp(ip string) MysqlOption {
	return func(options *MysqlOptions) {
		options.Ip = ip
	}
}

func DBName(name string) MysqlOption {
	return func(options *MysqlOptions) {
		options.DBName = name
	}
}

func InitOrm(opts ...MysqlOption) error {
	options := new(MysqlOptions)
	for _, opt := range opts {
		opt(options)
	}
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", options.Username, options.Password, options.Ip, options.DBName))
	if err != nil {
		fmt.Println("zgz-db-utils 初始化失败")
		return err
	}
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(20)
	Engine = engine
	fmt.Println("zgz-db-utils 初始化成功")
	return nil
}

// SetEngine 直接设置引擎
func SetEngine(engine *xorm.Engine) {
	Engine = engine
}
