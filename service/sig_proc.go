package service

import "os"

type SigHandler func() error
type SignalProc map[os.Signal]SigHandler

type SigSet struct {
	m map[os.Signal]SigHandler
}

func NewSigMap() *SigSet {
	s := new(SigSet)
	s.m = make(map[os.Signal]SigHandler)
	return s
}

func (set *SigSet) Register(s os.Signal, handler SigHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *SigSet) HasHandler(sig os.Signal) bool {
	if _, found := set.m[sig]; found {
		return true
	} else {
		return false
	}
}

func (set *SigSet) Handle(sig os.Signal) (err error) {
	return set.m[sig]()
}
