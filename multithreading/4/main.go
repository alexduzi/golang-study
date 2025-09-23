package main

func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever
	// se não insere nada no canal e tenta pegar algum valor, essa operação
	// irá bloquear a main thread e como está vazio irá dar um deadlock
}
