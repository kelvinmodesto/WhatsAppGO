package main

import (
    
    
    )

type Node struct {
	id int
	mensagem MSG
}

type MSG struct {
}

type Queue struct {
	node  []*Node
	size  int
	head, tail int
	count int
	idUser int
	userName string
}

type QList struct {
	filas []*Queue
	size  int
	count int
}

func Push() {
    
}

func Pop(){
    
}

func main() {

}
