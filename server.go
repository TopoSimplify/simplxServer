package main

import (
	"sync"
)

type Server struct {
	Address      string
	Mode         string
	Exit         chan struct{}
	ExitWg       *sync.WaitGroup
}

func NewServer(address string, mode string) *Server {
	var exit          = make(chan struct{})
	var exitWg        = &sync.WaitGroup{}
	exitWg.Add(0)

	var server = &Server{
		Address:      address,
		Mode:         mode,
		Exit:         exit,
		ExitWg:       exitWg,
	}

	return server
}
