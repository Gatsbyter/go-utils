package service

import (
	"os"
	"os/signal"
)

var signalNotify = signal.Notify

type Service interface {
	Init() error

	Start() error

	Stop() error
}

func Run(service Service, proc SignalProc) error {
	if err := service.Init(); err != nil {
		return err
	}

	if err := service.Start(); err != nil {
		return err
	}

	sigHandler := NewSigMap()
	for sig, handler := range proc {
		sigHandler.Register(sig, handler)
	}

	signalChan := make(chan os.Signal, 10)
	signalNotify(signalChan)

	for {
		sig := <-signalChan

		if sigHandler.HasHandler(sig) {
			return sigHandler.Handle(sig)
		} else {
			return service.Stop()
		}
	}
}
