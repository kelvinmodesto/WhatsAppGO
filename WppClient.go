package main

import (
	"fmt"
	"strings"
)

//Região de variáveis globais
var destino string
var comando string
var msgAtual string

//Estruturas para tratamento de mensagens
func identificarComando(txt []string) {

}

//
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

func open() {

}

func close() {

}

func all() {

}

func send() {

}

func lerDestino() {

}

func lerMSG() {

}

func getListaContatos() {

}

//Estruturas para envio de mensagens via TCP

func main() {
	lerTexto("fdfds @Kelvin Oi, tudo bem Kelvin?")

}
