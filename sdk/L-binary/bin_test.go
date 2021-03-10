package L_binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"testing"
)

// binary包实现了对数据与byte之间的转换，以及varint的编解码

// 读二进制到变量里
func TestReadBin(t *testing.T) {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		panic(err)
	}

	fmt.Println(pi)
}

// 写变量到二进制形式
func TestWriteBin(t *testing.T) {
	buf := new(bytes.Buffer)
	pi := math.Pi

	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.Bytes())
}

func TestLittle(t *testing.T) {
	a := binary.LittleEndian.Uint16([]byte{3, 4})
	fmt.Printf("%x\n", a)
}

func TestBig(t *testing.T) {
	a := binary.BigEndian.Uint16([]byte{3, 4})
	fmt.Printf("%x\n", a)
}

// 判断大小端
func TestCheck(t *testing.T) {
	var a uint16 = 0x0102
	var b = uint8(a)

	if b == 0x02 {
		fmt.Println("LittleEndian")
	} else if b == 0x01 {
		fmt.Println("BigEndian")
	} else {
		fmt.Println("Fuck")
	}
}
