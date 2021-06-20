package initialize

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"process_data/common"
	"process_data/config"
	"runtime/debug"
	"time"
)

func InitSqlx() {
	flowDbConf := config.GConfig.DBS[0].Conf
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", flowDbConf.UserName, flowDbConf.Password, flowDbConf.Host, flowDbConf.Port, flowDbConf.Name))
	if err != nil {
		log.Fatal(fmt.Sprintf("mysql session 启动失败: %v\n堆栈信息: %v", err, string(debug.Stack())))
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf(fmt.Sprintf("mysql session  启动失败: %v\n堆栈信息: %v", err, string(debug.Stack())))
		os.Exit(0)
		return
	}

	common.Mysqlx = db

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Second * 300)
	log.Printf("InitSqlx Success.")
}
