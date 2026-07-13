package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"relatorio/conector"
	"relatorio/correlacionador"
	"relatorio/sumarizador"
	"sync"
	"time"
)

type RespostaAPI struct {
	EndpointName string
	Data         []byte
	Error        error
}

func carregarDadosDaAPI(nome string, url string, wg *sync.WaitGroup, ch chan<- RespostaAPI) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- RespostaAPI{EndpointName: nome, Error: err}
		return
	}
	defer resp.Body.Close()

	corpoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- RespostaAPI{EndpointName: nome, Error: err}
		return
	}

	ch <- RespostaAPI{EndpointName: nome, Data: corpoBytes}
}

func carregarDadosSequencial(configEnv conector.ConfigEnvironment) conector.NormRecords {
	registrosSequenciais := make(conector.NormRecords)
	for nome, value := range configEnv {
		endpointConfig := value.(conector.ConfigEndpoint)
		urlDinâmica := endpointConfig.EndpointLocation
		if urlDinâmica == "" {
			continue
		}

		resp, err := http.Get(urlDinâmica)
		if err != nil {
			continue
		}
		corpoBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		registroNormalizado, err := conector.NormalizeData(resp.Data, endpointConfig, resp.EndpointName)
		if err != nil {
			fmt.Printf("[Erro] Falha ao normalizar dados de %q: %v\n", resp.EndpointName, err)
			continue
		}
		registrosSequenciais[nome] = registroNormalizado
	}
	return registrosSequenciais
}

func main() {
	// -------------------------------------------------------------------------
	// ETAPA 1: Carregar Configurações YAML Dinamicamente
	// -------------------------------------------------------------------------
	// 1. Descobre o caminho absoluto do executável atual
	execPath, err := os.Executable()
	if err != nil {
	    fmt.Printf("[Erro] Não foi possível encontrar o caminho do executável: %v\n", err)
	    os.Exit(1)
	}

	// 2. Pega a pasta onde o executável está localizado
	execDir := filepath.Dir(execPath)

	// 3. Monta o caminho apontando para a pasta "APIs" ao lado do executável
	apisPath := filepath.Join(execDir, "APIs")

	// 4. Cria o sistema de arquivos apontando para o novo caminho relativo dinâmico
	dirFS := os.DirFS(apisPath)
	
	configEnv, errs := conector.LoadConfigFiles(dirFS, ".")
	if len(errs) > 0 {
		fmt.Printf("[Erro] Falha ao carregar configurações: %v\n", errs)
		os.Exit(1)
	}

	rootTable, errs := conector.CompileConfigEnvironment(configEnv)
	if len(errs) > 0 {
		fmt.Printf("[Erro] Falha ao compilar ambiente: %v\n", errs)
		os.Exit(1)
	}
	fmt.Printf("[Info] Tabela Raiz (Nó Pai): %q\n", rootTable)

	// -------------------------------------------------------------------------
	// ETAPA 2: Disparar Requisições Concorrentes usando URLs do YAML
	// -------------------------------------------------------------------------
	// -------------------------------------------------------------------------
	// DEBUGGING DE DESEMPENHO: CONCORRENTE VS SEQUENCIAL
	// -------------------------------------------------------------------------
	fmt.Println("\n[Debug] Iniciando Teste Comparativo de Coleta...")

	// Medição 1: Abordagem Sequencial
	inicioSeq := time.Now()
	_ = carregarDadosSequencial(configEnv) // Apenas simula para pegar o tempo
	duracaoSeq := time.Since(inicioSeq)

	// Medição 2: Abordagem Concorrente Real (O código que já tínhamos)
	inicioPar := time.Now()
	var wg sync.WaitGroup
	canalRespostas := make(chan RespostaAPI, len(configEnv))

	for nome, value := range configEnv {
		endpointConfig := value.(conector.ConfigEndpoint)
		urlDinâmica := endpointConfig.EndpointLocation
		if urlDinâmica == "" {
			continue
		}
		wg.Add(1)
		go carregarDadosDaAPI(nome, urlDinâmica, &wg, canalRespostas)
	}

	go func() {
		wg.Wait()
		close(canalRespostas)
	}()

	// -------------------------------------------------------------------------
	// ETAPA 3: Coletar e Normalizar os Dados (Módulo Conector)
	// -------------------------------------------------------------------------
	todosOsRegistros := make(conector.NormRecords)
	for resp := range canalRespostas {
		if resp.Error != nil {
			fmt.Printf("[Erro] Falha ao consumir %q: %v\n", resp.EndpointName, resp.Error)
			continue
		}
		endpointConfig := configEnv[resp.EndpointName].(conector.ConfigEndpoint)
		registroNormalizado, err := conector.NormalizeData(resp.Data, endpointConfig, resp.EndpointName)
		if err != nil {
			fmt.Printf("[Erro] Falha ao normalizar dados de %q: %v\n", resp.EndpointName, err)
			continue
		}
		
		todosOsRegistros[resp.EndpointName] = registroNormalizado
	}
	duracaoPar := time.Since(inicioPar)

	// Exibe o veredito no terminal
	fmt.Printf("\n====== VEREDITO DE PERFORMANCE ======\n")
	fmt.Printf("Tempo no modo Sequencial : %v\n", duracaoSeq)
	fmt.Printf("Tempo no modo Concorrente: %v\n", duracaoPar)
	fmt.Printf("Ganho de Desempenho      : %.2f vezes mais rápido!\n", float64(duracaoSeq)/float64(duracaoPar))
	fmt.Printf("=====================================\n\n")

	// -------------------------------------------------------------------------
	// ETAPA 4: Correlacionar e Unificar os Dados baseados nas FKs (Correlacionador)
	// -------------------------------------------------------------------------
	fmt.Println("\n[Info] Iniciando a correlação de Chaves Estrangeiras (Joins)...")
	
	// Executa a árvore de unificação a partir do nó raiz determinado dinamicamente
	tabelaUnificada := correlacionador.MakeUnifiedData(todosOsRegistros, configEnv, rootTable)
	
	fmt.Printf("[Sucesso] Dados unificados com base na raiz %q. Total de linhas: %d\n", rootTable, len(tabelaUnificada))

	// -------------------------------------------------------------------------
	// ETAPA 5: Converter e Sumarizar Métricas Analíticas (Sumarizador)
	// -------------------------------------------------------------------------
	fmt.Println("\n[Info] Iniciando processamento de métricas no Sumarizador...")

	// O sumarizador consome o tipo []JData (que possui a assinatura map[string]any)
	// Como seu normRecord é []normData (onde normData também é map[string]any),
	// podemos converter ou fazer a ponte de tipos aqui:
	dadosParaSumarizar := make([]sumarizador.JData, len(tabelaUnificada))
	for i, v := range tabelaUnificada {
		dadosParaSumarizar[i] = sumarizador.JData(v)
	}

	// Exemplo prático de métrica usando suas funções de sumarização:
	// Vamos agrupar os dados pela criticidade dos endpoints e contar o volume
	colunaAgrupamento := []string{"endpoints.endpoint_criticidade"}
	resultadoMetricas := sumarizador.Count(dadosParaSumarizar, colunaAgrupamento, "endpoints.endpoint_id")

	// Imprime os indicadores finais consolidados
	fmt.Println("\n================ RESULTADO DO RELATÓRIO OBTIDO ================")
	for _, metrica := range resultadoMetricas {
		fmt.Printf("Criticidade: %v | Total de Ocorrências: %v\n", 
			metrica["endpoints.endpoint_criticidade"], 
			metrica["count_"],
		)
	}
	fmt.Println("===============================================================")

	fmt.Printf("\n[OK] Pipeline executada com sucesso absoluto! Tempo total: %v\n", time.Since(inicioSeq))
}
