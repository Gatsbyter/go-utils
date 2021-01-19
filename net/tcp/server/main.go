package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Msg struct {
	MagicNum [4]byte
	Len [2]byte
	Data []byte
}

// 玩一玩粘包
func main()  {
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}

	defer lis.Close()
	log.Printf("listen to %s --- %s", lis.Addr().Network(), lis.Addr().String())

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		} else {
			go handle(conn)
		}
	}
}

func handle(conn net.Conn) {
	defer func() {
		log.Printf("close conn with %s", conn.RemoteAddr().String())
		conn.Close()
	}()

	log.Printf("new connection with %s", conn.RemoteAddr())
    result := bytes.NewBuffer(nil)
    i := 1
	
	for {
		buf := make([]byte, 65542)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				log.Print(err)
			}
		} else {
			result.Write(buf[:n])
			scanner := bufio.NewScanner(result)
			scanner.Split(packetSlitFunc)
			for scanner.Scan() {
				log.Printf("%d ---- recv: %s", i, string(scanner.Bytes()[6:]))
				i++
			}
		}
		result.Reset()
	}
}

func packetSlitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0x123456 {
		var l int16
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
		binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &l)
		pl := int(l) + 6
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}