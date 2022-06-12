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

	connection := makeMongoDbConnection()
	defer connection.Disconnect(context.Background())

	deps := createRestDeps(connection.Client.Database("OPD"))
	restServer := rest.Server{Deps: &deps}

	isInterrupted, ctx, cancelRoutines := makeRoutineController()

	restServer.Start("8080")
	go outbox.RunOutboxRelay(ctx)
	go worker.RunWorker(ctx)

	<-isInterrupted
	cancelRoutines()

	fmt.Println("exiting")
	if err := restServer.Stop(); err != nil {
		panic(err)
	}
}

func makeRoutineController() (chan os.Signal, context.Context, context.CancelFunc) {
	stop := makeStopSignal()
	child, cancel := context.WithCancel(context.Background())
	return stop, child, cancel
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
