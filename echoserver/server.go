package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/tominescu/double-golang/simplelog"
)

func perror(prefix string, err error) {
	fmt.Fprintf(os.Stderr, "%s:%s\n", prefix, err)
}

func handle_error(prefix string, err error) {
	if err != nil {
		perror(prefix, err)
		log.Info("Server stopped with code:1")
		os.Exit(1)
	}
}

func signal_handler(ch chan os.Signal) {
	sig := <-ch
	log.Warn("Received signal %d", sig)
	os.Exit(0)
	log.Info("Server stopped with code:0")
}

func main() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGINT, syscall.SIGTERM)
	go signal_handler(signal_chan)

	listen_port := 0
	port := flag.Int("p", 9900, "which port to bind")
	flag.Parse()
	if *port > 0 && *port < 65536 {
		listen_port = *port
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", listen_port))
	handle_error("listen", err)
	defer ln.Close()
	log.Info("Server listen on :::%d", listen_port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			perror("accept", err)
		}
		defer conn.Close()

		go HandleClientConn(conn)
	}
	log.Info("Server stopped with code:0")
}

func HandleClientConn(conn net.Conn) {
	log.Info("New client from %s", conn.RemoteAddr())
	n, _ := io.Copy(conn, conn)
	log.Info("Sent %d bytes to client %s", n, conn.RemoteAddr())
	log.Info("Close client %s", conn.RemoteAddr())
}
