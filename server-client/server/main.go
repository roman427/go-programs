package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func connHandler(conn *net.Conn) {

	for {
		msg, err := bufio.NewReader(*conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", msg)
	}

}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go connHandler(&conn)
	}

}
