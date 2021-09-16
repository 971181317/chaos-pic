package main

import "log"

func main() {
	InitLog("log")
	if err := RegisterRoute().Run(":8080"); err != nil {
		log.Fatalf("server start error %s", err.Error())
	}
	//读取控制台参数
	//args := os.Args
	////处理无参数情况
	//if len(args) == 1 {
	//	fmt.Print("-h 获取帮助")
	//} else {
	//	switch args[1] {
	//	case "-h":
	//		fmt.Println("-dn 获取设备信息")
	//		fmt.Println("-catch 实时抓包，保存内容到filePath")
	//		fmt.Println("-offline 从filePath中获取解析离线数据")
	//		fmt.Println("-speed 监控实时网速")
	//	case "-dn":
	//	default:
	//		fmt.Print("错误参数, -h获取帮助")
	//	}
	//}
	//fmt.Println("11")
	//imgData, _ := ioutil.ReadFile("/Users/bytedance/Desktop/650-686x240.jpeg")
	//buf := bytes.NewBuffer(imgData)
	//image, err := imaging.Decode(buf)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
