package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func acessarAPI(nome string, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	inicio := time.Now()
	fmt.Printf("[Info] Iniciando Chamada para %s...\n", nome)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[Erro] %s falhou: %v\n", nome, err)
		return
	}
	defer resp.Body.Close()
	
	corpoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[Erro] Falha ao ler o corpo da resposta de %s: %v\n", nome, err)
		return
	}

	duracao := time.Since(inicio)
	fmt.Printf("[Sucesso] %s respondeu com Status %s em %v\n", nome, resp.Status, duracao)
	
	fmt.Printf("[Info] Json de %s: %s\n\n", nome, string(corpoBytes))
}

func main() {
	var wg sync.WaitGroup

	api1 := "http://mock-api-1/usuarios/"
	api2 := "http://mock-api-2/endpoints/"
	api3 := "http://mock-api-3/subredes/"
	api4 := "http://mock-api-4/logs_acesso/"

	tempoTotal := time.Now()

	wg.Add(4)
	go acessarAPI("usuarios", api1, &wg)
	go acessarAPI("endpoints", api2, &wg)
	go acessarAPI("subredes", api3, &wg)
	go acessarAPI("logs_acesso", api4, &wg)

	fmt.Println("Aguardando as APIs responderem...")
	wg.Wait()

	fmt.Printf("\nProcesso concluído! Tempo total de execução: %v\n", time.Since(tempoTotal))
}
