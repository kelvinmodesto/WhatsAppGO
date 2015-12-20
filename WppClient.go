package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var buffer []byte
var CONN net.Conn

type Cliente struct {
	name string
	conn net.Conn
	quit bool
}

func (cl *Cliente) inicializarConversa() {
	cl.conn, _ = net.Dial("tcp", "kelvinmodesto.koding.io:9933")
	fmt.Println("OK!")
	buffer = []byte("@open")
	cl.conn.Write(buffer)
}

//Região de variáveis globais

func (cl *Cliente) send() {
	cl.conn, _ = net.Dial("tcp", "localhost:9933")
	fmt.Println("OK!")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Eu: ")
		msg, _ := reader.ReadString('\n')
		buffer = []byte(msg)
		cl.conn.Write(buffer)
	}
}

func (cl *Cliente) getINBOX() {

}

func (cl *Cliente) receiveMessage() {
	ln, _ := net.Listen("tcp", ":9933")
	cl.conn, _ = ln.Accept()
	reader := bufio.NewReader(cl.conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}

}

func (cl *Cliente) fecharConexao() {
	cl.quit = true
}

func main() {
	clienteTeste := &Cliente{"@kelvin", CONN, false}
	clienteTeste.send()
}
