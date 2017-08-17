package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/tominescu/double-golang/simplelog"
)

func main() {
	port := flag.Int("p", 9999, "Listen port")
	flag.Parse()
	if *port > 65535 || *port <= 0 {
		*port = 9999
	}
	portStr := fmt.Sprintf(":%d", *port)

	// udp
	go ListenUDP(*port)

	// tcp
	ln, err := net.Listen("tcp", portStr)
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	buf := fmt.Sprintf("%s\n", remoteAddr.String())
	simplelog.Info("Recv Client %s", remoteAddr.String())

	w := bufio.NewWriter(conn)
	w.WriteString("HTTP/1.1 200 OK\r\n")
	w.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	lengthStr := fmt.Sprintf("Content-Length: %d\r\n", len(buf))
	w.WriteString(lengthStr)
	w.WriteString("\r\n")
	w.WriteString(buf)
	w.Flush()

	conn.Close()
}

func ListenUDP(port int) {
	var err error
	ln, err := net.ListenUDP("udp", &net.UDPAddr{Port: port})
	if err != nil {
		fmt.Printf("Listen udp error: %s\n", err)
		return
	}
	for {
		buf := make([]byte, 1024)
		n, remoteAddr, err := ln.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("Read udp error: %s\n", err)
			continue
		}
		simplelog.Info("Read %d bytes from udp client[%s:%d]: %s", n, remoteAddr.IP, remoteAddr.Port, strings.TrimSpace(string(buf[:n])))
		_, err = ln.WriteToUDP([]byte(fmt.Sprintf("%s:%d\n", remoteAddr.IP, remoteAddr.Port)), remoteAddr)
		if err != nil {
			fmt.Printf("Write udp to client[%s:%d] error: %s\n", remoteAddr.IP, remoteAddr.Port, err)
		}
	}
}
