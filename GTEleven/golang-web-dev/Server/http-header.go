// https://www.youtube.com/watch?v=x9fhKK0IcME&index=40&list=PLSak_q1UXfPpXj-q1BeucvBAlNdotQWVD

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	li, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(conn)
	}
}
