package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func perror(prefix string, err error) {
	fmt.Fprintf(os.Stderr, "%s:%s", prefix, err)
}

func handle_error(prefix string, err error) {
	if err != nil {
		perror(prefix, err)
		os.Exit(1)
	}
}

func usage(name string) {
	fmt.Fprintf(os.Stderr, "%s -p <remote_port> remote_addr\n")
	flag.PrintDefaults()
}

func main() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGTERM, syscall.SIGINT)
	go func(ch chan os.Signal) {
		<-ch
		os.Exit(0)
	}(signal_chan)

	port := flag.Int("p", 9900, "remote port")
	flag.Parse()

	remote_port := 0
	if *port > 0 && *port < 65536 {
		remote_port = *port
	}

	remote_addr := "localhost"
	if flag.NArg() == 1 {
		remote_addr = flag.Arg(0)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", remote_addr, remote_port))
	handle_error("dial", err)
	defer conn.Close()

	fmt.Fprintf(os.Stderr, "Connected to %s:%d\n", remote_addr, remote_port)
	fmt.Fprintf(os.Stderr, "Press Ctrl + C to quit\n")
	go RecvConn(conn)

	sr := bufio.NewReader(os.Stdin)
	for {
		_, err := sr.WriteTo(conn)
		if err != nil {
			break
		}
	}
}

func RecvConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		nread, err := conn.Read(buf)
		if err != nil {
			os.Exit(0)
		}
		fmt.Printf("%s", string(buf[0:nread]))
	}
}
