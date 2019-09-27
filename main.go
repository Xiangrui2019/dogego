package main

import (
	_ "dogego/conf"
	"dogego/protocol"
	"dogego/routers"
)

func main() {
	router := routers.NewRouter()
	channel := make(chan bool)

	// 启动所有的服务, 也就是所有注册进去的Protocol
	protocol.StartAllServers(router)

	<-channel
}
