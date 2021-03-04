package L_sync

import "sync"

// 使用channel 来实现互斥
var (
	// 为了同步的话 可以这样用struct{} 省内存
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	mu      sync.Mutex               // guards balance
	balance int
)

func DepositSema(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func BalanceSema() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

/* ------------------------------------------------------------------------------------------------- */

func DepositMu(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func BalanceMu() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

// 这个函数会有很严重的问题
// 重入锁
func WithdrawMu(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	DepositMu(-amount)
	if BalanceMu() < 0 {
		DepositMu(amount)
		return false
	}
	return true
}
