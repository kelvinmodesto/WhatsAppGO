package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//Região de variáveis globais

func send() {
	conn, _ := net.Dial("tcp", "kelvinmodesto.koding.io:9933")
	fmt.Println("OK!")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Eu: ")
		msg, _ := reader.ReadString('\n')
		buff := []byte(msg)
		conn.Write(buff)
	}
}

func main() {
}
