package app

import (
	"fmt"
	"os"
	"os/signal"
	"queueing-clean-demo/app/deps"
	"queueing-clean-demo/base"
	"queueing-clean-demo/outbox"
	"queueing-clean-demo/rest"
	"queueing-clean-demo/worker"
	"syscall"
)

func StartApp() {
	SetupMessageBroker()

	servers := []base.IServer{
		outbox.NewServer(deps.NewOutboxDeps, "allEvents"),
		worker.NewServer(deps.NewWorkerDeps, "allEvents"),
		rest.NewServer(deps.NewRestDeps, "8080"),
	}

	isInterrupted := makeStopSignal()

	startAllServers(servers)

	<-isInterrupted
	fmt.Println("exiting")

	stopAllServersDescending(servers)
}

func startAllServers(servers []base.IServer) {
	for _, server := range servers {
		server.Start()
	}
}

func stopAllServersDescending(servers []base.IServer) {
	for i := len(servers) - 1; i >= 0; i-- {
		if err := servers[i].Stop(); err != nil {
			println(err.Error())
		}
	}
}

func makeStopSignal() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	return stop
}
