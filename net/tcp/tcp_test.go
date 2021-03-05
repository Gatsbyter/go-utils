package tcp

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

// ⚠️ 有一个不错的工具 叫nc

// 一个简单的指定本地端口的telnet工具
func TestTcp(t *testing.T) {
	localPort := 56580

	netAddr := &net.TCPAddr{Port: localPort}

	dialer := net.Dialer{Timeout: time.Second * 10, LocalAddr: netAddr}

	conn, err := dialer.Dial("tcp", "148.70.43.160:9995")
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}

	defer conn.Close()

	time.Sleep(time.Hour)
}
