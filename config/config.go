package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ListenConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type InnerListenConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DBConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type DBConfEntry struct {
	DbName string `json:"db_name"`
	Conf   DBConf `json:"conf"`
}

type RedisConf struct {
	Prefix   string `json:"prefix"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"initialize"`
	MaxConn  int    `json:"max_conn"`
}

type LogConf struct {
	LogLevel    string `json:"loglevel"`
	LogFileDir  string `json:"logfile_dir"`
	LogFileName string `json:"logfile_name"`
}

type Configuration struct {
	Listen      ListenConf        `json:"listen"`
	InnerListen InnerListenConf   `json:"inner_listen"`
	ApiPrefix   string            `json:"api_prefix"`
	CenterApi   map[string]string `json:"center_api"`
	DBS         []DBConfEntry     `json:"db_list"`
	Redis       RedisConf         `json:"redis"`
	Log         LogConf           `json:"mock_data"`
	Kafka       KafkaConf         `json:"kafka"`
	RPC         RpcConf           `json:"rpc"`
}

type KafkaConf struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
	GroupId string   `json:"group_id"`
}

type RpcConf struct {
	AdminCenter string `json:"admin_center"`
	UserCenter  string `json:"user_center"`
}

var GConfig = Configuration{}

func (self *Configuration) Init() {
	// 读取配置文件
	absPath, _ := os.Getwd()
	fp, err := os.Open(fmt.Sprintf("%s/default.cfg", absPath))
	if err != nil {
		log.Fatal(fmt.Sprintf("config init err: %v", err))
		return
	}
	defer fp.Close()
	// 解析为Json
	if err := json.NewDecoder(fp).Decode(&GConfig); err != nil {
		log.Printf("config init NewDecoder err: %v", err)
		return
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//mock_data.Printf("----debug----Configuration: \n%+v\n", GConfig)
}

func Init() {
	GConfig.Init()
}
