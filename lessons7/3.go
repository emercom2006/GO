package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {

		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))

		var exit string
		fmt.Println("Введите exit для выхода")
		fmt.Scanln(&exit)
		if exit == "exit" {
			os.Exit(1)
		}
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

	}
}
