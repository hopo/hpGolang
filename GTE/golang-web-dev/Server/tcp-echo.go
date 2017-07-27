// https://ko.wikipedia.org/wiki/%EC%A0%84%EC%86%A1_%EC%A0%9C%EC%96%B4_%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C

package main

import (
	"io"
	"net"
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

		// handels only one connection
		io.Copy(conn, conn)
		conn.Close()
	}
}
