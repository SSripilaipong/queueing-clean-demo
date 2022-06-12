package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"queueing-clean-demo/outbox"
	"queueing-clean-demo/rest"
	"queueing-clean-demo/toolbox/mongodb"
	"queueing-clean-demo/worker"
	"syscall"
)

func StartApp() {
	SetupMessageBroker()

	restServer := rest.NewServer(newRestDeps, "8080")
	outboxServer := outbox.NewServer()
	workerServer := worker.NewServer()

	isInterrupted := makeStopSignal()

	restServer.Start()
	outboxServer.Start()
	workerServer.Start()

	<-isInterrupted

	fmt.Println("exiting")
	if err := restServer.Stop(); err != nil {
		println(err.Error())
	}
	if err := outboxServer.Stop(); err != nil {
		println(err.Error())
	}
	if err := workerServer.Stop(); err != nil {
		println(err.Error())
	}
}

func makeMongoDbConnection() *mongodb.Connection {
	connection, err := mongodb.CreateConnection(context.Background(), "root", "admin", "mongodb", "27017")
	if err != nil {
		panic(err)
	}
	return connection
}

func makeStopSignal() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	return stop
}
