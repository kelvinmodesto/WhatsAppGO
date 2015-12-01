package main

type Node struct {
	mensagem MSG
}

type MSG struct {
}

type Queue struct {
	node  []*MSG
	size  int
	head  int
	tail  int
	count int
}

type QList struct {
	filas []*Queue
	size  int
	count int
}

type Receptor struct {
}

func add() {

}

func remove() {

}

func send() {

}

func receive() {

}

func listen() {

}

func main() {

}
