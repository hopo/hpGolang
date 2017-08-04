// https://www.youtube.com/watch?v=dd1sK2g1_-Q&t=1s&index=33&list=PLSak_q1UXfPpXj-q1BeucvBAlNdotQWVD#t=453.265124

// what is tcp?
// https://ko.wikipedia.org/wiki/%EC%A0%84%EC%86%A1_%EC%A0%9C%EC%96%B4_%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("FROM SERVER: HELLO CLASS ", time.Now(), "\n"))

		conn.Close()
	}
}

// $ telnet localhost 4000
