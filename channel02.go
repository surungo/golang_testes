package main

import (
	"fmt"
	"time"
)

// Função que simula o envio de um número após algum cálculo ou operação.
// Ela envia um número para o canal fornecido.
func sendNumber(ch chan<- int, num int, time_v int) {

	time.Sleep(time.Duration(time_v) * time.Second)
	var originalValue int
	//	originalValue = -> ch
	originalValue = <-ch
	fmt.Println("Valor original no canal antes de alterar:", originalValue, time_v)

	//var originalValue int
	// select {
	// case originalValue = <-ch:
	// 	fmt.Println("Valor original no canal antes de alterar:", originalValue)
	// default:
	// 	fmt.Println("Canal vazio, nenhum valor original encontrado.")
	// }
	ch <- num // Envia 'num' através do canal.
}

func main() {
	// Cria um canal de inteiros.
	numChannel := make(chan int)

	// Inicia a goroutine sendNumber para enviar um número através do canal.
	go sendNumber(numChannel, 20, 2)
	go sendNumber(numChannel, 40, 4)
	go sendNumber(numChannel, 60, 6)
	go sendNumber(numChannel, 80, 8)

	time.Sleep(3 * time.Second)
	num := <-numChannel
	fmt.Println("Recebi o número:", num)
	time.Sleep(2 * time.Second)
	num = <-numChannel
	fmt.Println("Recebi o número:", num)
	time.Sleep(2 * time.Second)
	num = <-numChannel
	fmt.Println("Recebi o número:", num)
	time.Sleep(2 * time.Second)
	num = <-numChannel
	fmt.Println("Recebi o número:", num)
}
