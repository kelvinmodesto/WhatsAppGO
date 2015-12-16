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
	sender string
	text   string
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

// func addNewUser(user *User) {
// 	userMap := make(map[user.username]user)
// 	if usuario, ok := userMap[user.username]; ok{
// 		user.inbox.PushQueue(n)

// 		}
// 		else{

// 		}
// }

//Operações com Usuários
func logon() {

}

func logoff() {

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

func receive() {
	fmt.Println("Starting...")
	ln, _ := net.Listen("tcp", ":9933")
	fmt.Println("Listening...")
	for {
		conn, _ := ln.Accept()
		fmt.Println("New Connection!")
		go connection(conn)
	}

}
func connection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		//fmt.Println(msg)
		var novaMSG []string
		novaMSG = lerTexto(msg)
		texto := &MSG{novaMSG[0], novaMSG[1]}
		node := &Node{contadorNode, texto}
		contadorNode++
		adicionarMSG(novaMSG[0], node)
	}
}

func adicionarMSG(username string, texto *Node) {
	if usuario, ok := userMap[username]; ok {
		usuario.inbox.PushQueue(texto)
	} else {
		newUser := User{contadorUser, username, createQueue(1, username), true}
		userMap[username] = newUser
		userMap[username].inbox.PushQueue(texto)
	}
}

func inicializarUserMap() {
	userMap = make(map[string]User)
}

func main() {
	inicializarUserMap()
	var novaMSG []string
	novaMSG = lerTexto("@kelvin testando codigo")
	texto := &MSG{novaMSG[0], novaMSG[1]}
	node := &Node{contadorNode, texto}
	contadorNode++
	adicionarMSG(novaMSG[0], node)
	//fmt.Println(userMap["@kelvin"])
	fmt.Println(userMap["@kelvin"].inbox.nodes[0].mensagem.text)
}
