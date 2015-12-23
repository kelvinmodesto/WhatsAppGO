package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
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

func receive() {
	ln, _ := net.Listen("tcp", ":9933")
	var conn net.Conn
	for {
		conn, _ = ln.Accept()
		go connection(conn, ln)
		//go sendControl()
	}

}

func sendControl() {
	msgToSend := getINBOX()
	for {
		defer time.Sleep(5000 * time.Millisecond)
		for i := 0; i < len(msgToSend); i++ {
			sendTo(msgToSend[i].mensagem, userMap[msgToSend[i].mensagem.receiver].address)
		}
	}
}

func sendTo(mensagem *MSG, address string) {
	addressNovaPorta := strings.Split(address, ":")
	conn, _ := net.Dial("tcp", addressNovaPorta[0]+":9933")
	saida := mensagem.sender + ":" + mensagem.text
	for {
		buffer := []byte(saida)
		conn.Write(buffer)
	}
}

func connection(conn net.Conn, ln net.Listener) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		msgCompleta := lerSender(msg)
		var novaMSG []string
		novaMSG = lerTexto(msgCompleta[1])
		fmt.Println(novaMSG[0])
		destinatario := strings.Replace(novaMSG[0], "\n", "", 1)
		if destinatario == "@close" {
			fmt.Println("Desconectando...")
			userTemp := userMap[destinatario]
			userTemp.online = false
			userMap[destinatario] = userTemp
			conn.Close()
		} else {
			enderecoOrigem := conn.RemoteAddr().String()
			userTemp := userMap[destinatario]
			userTemp.address = enderecoOrigem
			userTemp.online = true
			userMap[destinatario] = userTemp

			texto := &MSG{novaMSG[0], msgCompleta[0], novaMSG[1]}
			node := &Node{contadorNode, texto}
			contadorNode++
			adicionarMSG(novaMSG[0], node)
		}

	}
}
func adicionarMSG(username string, texto *Node) {
	if usuario, ok := userMap[username]; ok {
		usuario.inbox.PushQueue(texto)
	} else {
		fmt.Println("Usuário não está cadastrado")
	}
}

func getINBOX() []*Node {
	var vetor []*Node
	j := 0
	for username, _ := range userMap {
		if i := userMap[username].inbox.count; i > 0 && userMap[username].online == true {
			for i > 0 {
				//esvaziar a fila

				vetor[j] = userMap[username].inbox.PopQueue()
				i -= 1
				j += 1
			}
			return vetor
		}
	}
	return nil
}

func inicializarUserMap() {
	userMap = make(map[string]User)
}

func main() {
	inicializarUserMap()
	userMap["@kelvin"] = User{0, "@kelvin", createQueue(1, "@kelvin"), false, "nenhum"}
	userMap["@thiago"] = User{1, "@thiago", createQueue(1, "@thiago"), false, "nenhum"}
	fmt.Println("Servidor Online")
	receive()
}
