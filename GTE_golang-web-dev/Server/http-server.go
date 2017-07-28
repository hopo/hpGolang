package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := sanner.Text()
		fmt.Println(ln)

		if i == 0 {
			method := strings.Fields(ln)[0]
			fmt.Println("METHOD", method)
		} else {
			// in headers now
			// when line in empty, header is done
			if ln == "" {
				break
			}
		}
		i++
	}

	// response
	body := `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title></title>
    </head>
    <body>
        <strong>Adios World</strong>
    </body>
</html>
    `
	io.WriteString(conn, "HTTP/1.1 302 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "Location: http://www.google.com\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
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
