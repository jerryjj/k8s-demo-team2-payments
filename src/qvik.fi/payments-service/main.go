package main

import (
	"os"
	"os/signal"
	"syscall"

	logging "github.com/op/go-logging"
)

// Initializes our local logger
func SetupLocalLogger(logModule string) *logging.Logger {
	var format = logging.MustStringFormatter("%{color}%{time:15:04:05.000} " +
		"%{shortfunc} ▶ %{level} " +
		"%{color:reset} %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)

	log := logging.MustGetLogger(logModule)
	logging.SetLevel(logging.DEBUG, logModule)

	// Compensate for all the wrapping layers around the logger
	log.ExtraCalldepth = 3

	log.Debug("Debug logging enabled.")

	return log
}

var (
	log = SetupLocalLogger("payments")
)

func main() {
	// Run the gRPC server in a goroutine not to block the main one..
	grpcServer := mustRunGrpcServer(50051)

	// ..and wait for a exit signal in the main routine
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	log.Debugf("Got signal: %v", s)

	// Do a clean shutdown
	grpcServer.GracefulStop()
}
