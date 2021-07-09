package main

import (
	"backend/chat"
	"flag"
)

var (
	// 命令行参数解析来指定端口号，默认是8080
	port = flag.String("p", ":8080", "set port")
)

func init() {
	flag.Parse()
}

func main() {
	chat.Start(*port)
}
