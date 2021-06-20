package main

import (
	"fmt"
	"process_data/config"
	"process_data/controller"
	"process_data/initialize"
)

func main() {
	var (
		err error
	)
	config.Init()
	_ = initialize.InitRedis()

	//1.bufSize 分别为10，20，50，100，200，1024，5120
	//err := controller.SaveSingleData("./mock_data", 10)
	//if err != nil {
	//	fmt.Println("", err)
	//}
	//2.设置kv
	err = controller.SaveMultiData("./mock_data")
	if err != nil {
		fmt.Println("SaveMultiData:", err)
	}

	fmt.Println("数据处理结束")
}
