package main

import "fmt"
import "net"
import "bufio"

func main(){
    fmt.Println("Starting...")
    ln, _ := net.Listen("tcp", ":9933")
    fmt.Println("Listening...")
    for {
        conn, _ := ln.Accept()
        fmt.Println("New Connection!")
        go connection(conn)
    }
}

func connection(conn net.Conn){
    reader := bufio.NewReader(conn) 
    for{
        msg, _ := reader.ReadString('\n')
        fmt.Println(msg)
    }
}