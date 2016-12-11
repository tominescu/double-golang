package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
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
	port := flag.Int("p", 9900, "remote port")
	flag.Parse()

	remote_port := 0
	if *port > 0 && *port < 65536 {
		remote_port = *port
	}

	remote_addr := "localhost"
	/*
		fmt.Printf("flag.NArg() == %d\n", flag.NArg())
		for i := 0; i < flag.NArg(); i++ {
			fmt.Printf("flag.Arg(%d) == %s\n", i, flag.Arg(i))
		}
	*/
	if flag.NArg() == 1 {
		remote_addr = flag.Arg(0)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", remote_addr, remote_port))
	handle_error("dial", err)
	defer conn.Close()

	go RecvConn(conn)

	sr := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	for {
		_, err := sr.Read(buf)
		if err != nil {
			break
		}
		_, err = conn.Write(buf)
		if err != nil {
			break
		}
	}
}

func RecvConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			os.Exit(0)
		}
		fmt.Printf("%s", string(buf))
	}
}
