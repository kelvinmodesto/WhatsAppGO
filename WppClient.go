package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var buffer []byte
var Conn net.Conn

//Região de variáveis globais

func Send() {
	conn, _ := net.Dial("tcp", "kelvinmodesto.koding.io:9933")
	fmt.Println("Digite o seu nome de usuário!")
	reader := bufio.NewReader(os.Stdin)
	header, _ := reader.ReadString('\n')
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print(strings.Replace(header, "\n", "", 1) + ":")
		msg, _ := reader.ReadString('\n')
		buffer = []byte(strings.Replace(header, "\n", "", 1) + ":" + msg)
		conn.Write(buffer)
	}
}

func receiveMessage() {
	ln, _ := net.Listen("tcp", ":9933")
	Conn, _ = ln.Accept()
	reader := bufio.NewReader(Conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}

}

func fecharConexao(conn net.Conn) {
	conn.Close()
}

func main() {
	// executarCliente()
	Send()
	//clienteTeste := &Cliente{"@kelvin", Conn, false}
	//clienteTeste.send()
	// go receiveMessage()

}
