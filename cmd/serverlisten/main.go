package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Writer interface {
	Write([]byte) (int, error)
}

type myWriter struct {
	buf []byte
}

func (m *myWriter) Write(p []byte) (int, error) {
	m.buf = append(m.buf, p...)
	fmt.Println("string(m.buf[:])")
	fmt.Println(string(m.buf[:]))
	return len(p), nil
}

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	var mw myWriter
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			// io.Copy(c, c)
			io.Copy(&mw, c)
			fmt.Println("sb.Len()")

			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
