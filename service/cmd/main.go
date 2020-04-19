package main

import (
	"github.com/Gatsbyter/go-utils/service"
	"sync"
)

type program struct {
	once sync.Once
}

func main() {
	prg := &program{}
	if err := service.Run(prg, service.SignalProc{}); err != nil {
	}
}

func (p *program) Init() error {
	return nil
}

func (p *program) Start() error {
	go func() {
	}()
	return nil
}

func (p *program) Stop() error {
	p.once.Do(func() {
	})
	return nil
}
