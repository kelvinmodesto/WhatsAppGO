package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//Região de variáveis globais

//Estruturas para tratamento de mensagens
// func identificarComando(txt []string) {
// 	switch cmd := strings.Replace(txt[0], "@", "", 1); cmd {
// 	case "close":
// 		close()
// 	case "open":
// 		open()
// 	case "all":
// 		all()
// 	case "help":
// 		help()
// 	case "inbox":
// 		inbox()
// 	default:
// 		id(txt)
// 	}
// }

//

func id(txt []string) {

}

func open() {

}

func close() {

}

func inbox() {

}

func all() {

}

func send() {
	conn, _ := net.Dial("tcp", "kelvinmodesto.koding.io:9933")
	fmt.Println("OK!")
	header := "Eu: "
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Digite: ")
		msg, _ := reader.ReadString('\n')
		buff := []byte(header + msg)
		conn.Write(buff)
	}
}

func lerDestino() {

}

func lerMSG() {

}

func getListaContatos() {

}

//Estruturas para envio de mensagens via TCP

func main() {
	send()
}
