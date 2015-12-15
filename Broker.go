package main

import (
	"fmt"
)

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

type QList struct {
	userList []*User
	size     int
	count    int
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

//Operações na lista de contatos
func createList(size int) *QList {
	return &QList{
		userList: make([]*User, size),
		size:     size,
		count:    0,
	}
}

func (ql *QList) addNewUser(user *User) {
	//tam := len(ql.userList)

}

//Operações com Usuários
func logon() {

}

func logoff() {

}

func lerTexto(msg string) {

	var txt []string
	if possuiComando := strings.Contains(msg, "@"); possuiComando == true {
		if indexComando := strings.Index(msg, "@"); indexComando == 0 {
			txt = strings.SplitN(msg, " ", 2)
			identificarComando(txt)
		} else {
			txt = strings.SplitN(msg, "", indexComando+1)
			txt = strings.SplitN(txt[indexComando], " ", 2)
			identificarComando(txt)
		}
	} else {
		if destino != "" {

		} else {
			fmt.Println("Eu:" + msg)
		}
	}
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
		fmt.Println(msg)
	}
}

func main() {
	q := createQueue(1, "tiago")
	q.PushQueue(&Node{1, &MSG{"kelvin", "Oi"}})
	q.PushQueue(&Node{2, &MSG{"kelvin", "Tudo Bem?"}})
	q.PushQueue(&Node{3, &MSG{"kelvin", "Quanto tempo cara"}})
	fmt.Println(q.PopQueue().mensagem.text)
	fmt.Println(q.PopQueue().mensagem.text)
	fmt.Println(q.PopQueue().mensagem.text)
	// us := &User{1, "tiago", q, false}

	// fmt.Println(addUser(us))
}
