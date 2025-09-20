package main

import "fmt"

func main() {
	// 1. Criação de um canal com make()
	// O canal 'mensagens' é do tipo string e não tem buffer (sincronizado)
	mensagens := make(chan string)

	// 2. Inicia uma nova goroutine
	go func() {
		// 3. Envia um valor para o canal usando <-
		mensagens <- "ping" // A seta aponta para onde o dado está indo
	}()

	// 4. Recebe o valor do canal
	// A seta agora aponta para o canal, indicando uma leitura
	msg := <-mensagens
	fmt.Println(msg) // Saída: ping

	// Exemplo com mais de uma goroutine

	// Canal para sinalizar o término da tarefa (bidirecional)
	done := make(chan bool)

	go func() {
		// Em uma goroutine, executa uma tarefa (neste caso, imprime "Executando...").
		fmt.Println("Executando a tarefa...")
		// Envia um sinal para o canal 'done' para indicar que a tarefa foi concluída.
		done <- true
	}()

	// No main, espera o sinal de término.
	<-done // Espera que a goroutine escreva no canal 'done'.
	fmt.Println("Tarefa concluída.")
}
