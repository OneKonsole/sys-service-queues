package main

import okmq "github.com/OneKonsole/sys-queueing"

func main() {
	queues, exchange := okmq.Initialize("../ini.json")
	okmq.LaunchService("amqp://guest:guest@localhost:5672/", exchange.Raw, queues)
}
