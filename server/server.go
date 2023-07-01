package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type FileServer struct { }

func createFileName() string {
	

	fileName := time.Now().Format(time.UnixDate) + ".log"
	

	return fileName
}

func (fs *FileServer) init() {



	ln, err := net.Listen("tcp", ":3333")

	if err != nil {
		log.Fatal(err)
	}
	
	for {
	
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Connection from: %s\n", conn.LocalAddr().String())
		go fs.read(conn)

	}

}

func (fs *FileServer) read(conn net.Conn) {

	fileName := createFileName()

	file, err := os.Create(fileName)

	fileWriter := bufio.NewWriter(file)
	if err != nil {
		return 
	}
	for {
		
		n, err := io.Copy(fileWriter, conn)
		if n == 0 {
			os.Remove(fileName)
		}
		if err != nil {
			return 
		}
		conn.Close()
		return 
	}
}


func main() {
	server := &FileServer{}
	server.init()

}
