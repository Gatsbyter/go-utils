package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	//ok()
	sticky()
}

// 不粘包版本
func ok() {
	data := []byte("[这里才是一个完整的数据包]")

	l := len(data)
	fmt.Println(l)
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0x123456)
	lenNum := make([]byte, 2)
	binary.BigEndian.PutUint16(lenNum, uint16(l))
	packetBuf := bytes.NewBuffer(magicNum)
	packetBuf.Write(lenNum)
	packetBuf.Write(data)

	conn, err := net.DialTimeout("tcp", "localhost:4444", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(packetBuf.Bytes())
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
	conn.Close()
}

// 粘包版本
func sticky() {
	data := []byte("[这里才是一个完整的数据包]")
	conn, err := net.DialTimeout("tcp", "localhost:4444", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i <1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
