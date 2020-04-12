package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"runtime"
)

var Port int
var Host string

const DebugMode = 0
const ReleaseMode = 1
const Error = 500
const Success = 200

func init() {
	flag.IntVar(&Port, "port", 8000, "host port")
	flag.StringVar(&Host, "host", "localhost", "host address")
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	var server = NewServer("localhost:8000", gin.DebugMode)
	var c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("signaled exit ...")
		close(server.Exit)
		os.Exit(1)
	}()

	server.Run()
}
