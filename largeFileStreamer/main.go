package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct {
}

func (f *FileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go f.readLoop(conn)
	}
}

func (f *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.Bytes())
		fmt.Printf("recieved %d bytes over the network\n", n)
	}
}

func writeFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	binary.Write(conn, binary.LittleEndian, int64(size))
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("written %d bytes are written on network\n", n)
	return nil
}
func main() {
	go func() {
		time.Sleep(4 * time.Second)
		writeFile(20000)
	}()
	server := &FileServer{}
	server.start()
}
