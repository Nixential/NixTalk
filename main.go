package main

import (
	"fmt"
	"net"
)

const PORT = "6969"

var connections []net.Conn

func handleConnection(conn net.Conn) {

	connections = append(connections, conn)

	fmt.Printf("Local Address %s\n", conn.LocalAddr().String())
	fmt.Printf("Remote Address %s\n", conn.RemoteAddr().String())

	conn.Write([]byte("Hello welcome to the chat\n"))

	for _, connection := range connections {
		connection.Write([]byte("A new connection has joined the chat!\n"))
	}

	handleMessage(conn)
}

func handleMessage(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading data from connection, ", err)
		}
		// fmt.Printf("Data received: %s\n", data[:n])
		for _, connection := range connections {
			if conn.RemoteAddr() != connection.RemoteAddr() {
				connection.Write([]byte(fmt.Sprintf("%s: %s", conn.RemoteAddr().String(), []byte(data[:n]))))
			}
		}
	}
}

func main() {
	fmt.Println("Hello World!")
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		fmt.Printf("Could not listen to awesome PORT %s!", PORT)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		go handleConnection(conn)
	}

}
