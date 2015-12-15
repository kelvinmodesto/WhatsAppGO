package main

import (
	"fmt"
)

//Estruturas do Broker

type User struct {
	id       int
	username string
	inbox    *Queue
	online   bool
}

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

type QList struct {
	userList []*User
	size     int
	count    int
}

//Zona de funções

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
	q.count++
	q.tail = q.count - 1
}

func (q *Queue) PopQueue() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
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
	tam := len(ql.userList)

}

//Operações com Usuários
func logon() {

}

func logoff() {

}

func main() {
	q := createQueue(1, "tiago")
	q.PushQueue(&Node{1, &MSG{"kelvin", "Oi"}})
	q.PushQueue(&Node{2, &MSG{"kelvin", "Tudo Bem?"}})
	q.PushQueue(&Node{3, &MSG{"kelvin", "Quanto tempo cara"}})
	us := &User{1, "tiago", q, false}

	fmt.Println(addUser(us))
}
