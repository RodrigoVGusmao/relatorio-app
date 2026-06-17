package main

import (
	"fmt"
	"net/http"
	"io"
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

	api1 := "http://mock-api-1/empregados/"
	api2 := "http://mock-api-1/setores/"
	api3 := "http://mock-api-2/maquinas/"

	tempoTotal := time.Now()

	wg.Add(3)
	go acessarAPI("empregados", api1, &wg)
	go acessarAPI("setores", api2, &wg)
	go acessarAPI("maquinas", api3, &wg)

	fmt.Println("Aguardando as APIs responderem...")
	wg.Wait()

	fmt.Printf("\nProcesso concluído! Tempo total de execução: %v\n", time.Since(tempoTotal))
}
