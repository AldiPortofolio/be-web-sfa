package main

import (
	"fmt"
	"os"
	"os/signal"
	"ottosfa-api-web/routers"
	"runtime"
	"strconv"
	"syscall"

	"go.uber.org/zap"
	"ottodigital.id/library/logger"
	"ottodigital.id/library/utils"
)

func main() {
	// init Logger Zap with file
	sugarLogger := logger.GetLogger()
	defer sugarLogger.Sync()

	maxProc, _ := strconv.Atoi(utils.GetEnv("MAXPROCS", "1"))

	runtime.GOMAXPROCS(maxProc)

	var errChan = make(chan error, 1)

	go func() {
		listenAddress := utils.GetEnv("OTTOSFA_API_WEB_PORT", "0.0.0.0:8045")

		fmt.Println("Starting @", listenAddress)
		errChan <- routers.Server(listenAddress)

	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		fmt.Println("got an interrupt, exiting...")
		sugarLogger.Error("Got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			fmt.Println("error while running api, exiting...", err)
			sugarLogger.Error("Error while running api, exiting... ", zap.Error(err))
		}
	}
}
