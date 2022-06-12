package app

import (
	"fmt"
	"os"
	"os/signal"
	"queueing-clean-demo/outbox"
	"queueing-clean-demo/rest"
	"queueing-clean-demo/worker"
	"syscall"
)

func StartApp() {
	SetupMessageBroker()

	outboxServer := outbox.NewServer()
	workerServer := worker.NewServer(newWorkerDeps, "allEvents")
	restServer := rest.NewServer(newRestDeps, "8080")

	isInterrupted := makeStopSignal()

	outboxServer.Start()
	workerServer.Start()
	restServer.Start()

	<-isInterrupted

	fmt.Println("exiting")
	if err := restServer.Stop(); err != nil {
		println(err.Error())
	}
	if err := workerServer.Stop(); err != nil {
		println(err.Error())
	}
	if err := outboxServer.Stop(); err != nil {
		println(err.Error())
	}
}

func makeStopSignal() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	return stop
}
