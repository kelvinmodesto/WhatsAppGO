package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var destino string
var comando string
var msgAtual string
var userMap map[string]User
var contadorNode int
var contadorUser int

//----------------Zona de Estruturas-------------

type Node struct {
	id       int
	mensagem *MSG
}

type MSG struct {
	sender   string
	receiver string
	text     string
}

type Queue struct {
	nodes    []*Node
	size     int
	head     int
	tail     int
	count    int
	idUser   int
	userName string
}

type User struct {
	id       int
	username string
	inbox    *Queue
	online   bool
	address  string
}

//--------------Zona de funções-----------------

//Funções do Broker

//Operações na fila
func createQueue(size int, userName string) *Queue {
	return &Queue{
		nodes:    make([]*Node, size),
		size:     size,
		userName: userName,
		count:    0,
	}

}

func (q *Queue) PushQueue(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) PopQueue() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func mudarStatus(user *User) {
	if user.online {
		user.online = false
	} else {
		user.online = true
	}
}

func lerSender(msg string) []string {
	return strings.Split(msg, ":")
}

func lerTexto(msg string) []string {

	var txt []string
	if possuiComando := strings.Contains(msg, "@"); possuiComando == true {
		if indexComando := strings.Index(msg, "@"); indexComando == 0 {
			txt = strings.SplitN(msg, " ", 2)
		} else {
			txt = strings.SplitN(msg, "", indexComando+1)
			txt = strings.SplitN(txt[indexComando], " ", 2)

		}
	}
	return txt
}

func send(mensagem *MSG, address string) {
	conn, _ := net.Dial("tcp", address)
	for {
		buffer := []byte(mensagem.sender + ": " + mensagem.text)
		conn.Write(buffer)
	}
}

func receive() {
	ln, _ := net.Listen("tcp", ":9933")
	conn, _ := ln.Accept()
	for {
		go connection(conn)
	}

}

func connection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		msgCompleta := lerSender(msg)
		var novaMSG []string
		novaMSG = lerTexto(msgCompleta[1])
		//switch novaMSG[0] {
		//case "@open":
		enderecoOrigem := conn.RemoteAddr().String()
		userTemp := userMap[novaMSG[0]]
		userTemp.address = enderecoOrigem
		userMap[novaMSG[0]] = userTemp

		//case "@close":
		checkClose(novaMSG[0], conn)

		//default:
		texto := &MSG{novaMSG[0], msgCompleta[0], novaMSG[1]}
		node := &Node{contadorNode, texto}
		contadorNode++
		adicionarMSG(novaMSG[0], node)
		fmt.Println(novaMSG[1])

		//}
	}
}
func adicionarMSG(username string, texto *Node) {
	if usuario, ok := userMap[username]; ok {
		usuario.inbox.PushQueue(texto)
		sendMSGToClient(usuario)
	} else {
		newUser := User{contadorUser, username, createQueue(1, username), true, "localhost"}
		userMap[username] = newUser
		userMap[username].inbox.PushQueue(texto)
		sendMSGToClient(newUser)
	}
}

func sendMSGToClient(user User) {
	for i := 0; i < user.inbox.count; i++ {
		send(&MSG{user.inbox.PopQueue().mensagem.sender, user.inbox.PopQueue().mensagem.receiver, user.inbox.PopQueue().mensagem.text}, user.address)
	}
}

func checkClose(cmd string, con net.Conn) {
	if cmd == "@close" {
		con.Close()
	}
}

func executeServer() {

}

func inicializarUserMap() {
	userMap = make(map[string]User)
}

func main() {
	inicializarUserMap()
	fmt.Println("Servidor Online")
	receive()
}
