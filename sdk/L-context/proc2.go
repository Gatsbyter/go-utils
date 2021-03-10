package L_context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type P2 struct {
	sync.WaitGroup
}

func (p *P2) StartProc(c context.Context) (context.CancelFunc, error) {
	ctx, cf := context.WithCancel(c)

	p.Add(1)

	go func(ctx context.Context) {
		defer func() {
			p.Done() // 等价于 p.Add(-1)
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("I've got cancel p2")
				return
			default:
				fmt.Println("I'm doing p2")
				time.Sleep(time.Second * 3)
			}
		}
	}(ctx)

	return cf, nil
}

func (p *P2) Close() {
	fmt.Println("I'm closing p2")
	p.Wait() //  WaitGroup到0 这个才会返回 也就是已经全部done掉了
	fmt.Println("p2 dead")
}
