package L_context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type P1 struct {
	// WaitGroup是routine安全的
	// WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址
	sync.WaitGroup
}

func (p *P1) StartProc(c context.Context) (context.CancelFunc, error) {
	ctx, cf := context.WithCancel(c)

	p.Add(1)

	go func(ctx context.Context) {
		defer func() {
			p.Done() // 等价于 p.Add(-1)
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("I've got cancel p1")
				return
			default:
				fmt.Println("I'm doing p1")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)

	return cf, nil
}

func (p *P1) Close() {
	fmt.Println("I'm closing p1")
	p.Wait() //  WaitGroup到0 这个才会返回 也就是已经全部done掉了
	fmt.Println("p1 dead")
}
