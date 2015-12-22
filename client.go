package main

import "fmt"
import "net"
import "os"
import "bufio"

func main() {
	fmt.Print("Connecting... ")
	conn, _ := net.Dial("tcp", "kelvinmodesto.koding.io:9933")
	fmt.Println("OK!")
	header := "Leonardo: "
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Digite: ")
		msg, _ := reader.ReadString('\n')
		buff := []byte(header + msg)
		conn.Write(buff)
	}
}
