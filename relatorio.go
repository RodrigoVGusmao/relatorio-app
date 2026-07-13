package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"relatorio/conector"
	"relatorio/correlacionador"
	"relatorio/gerador"      // Importado o novo pacote gerador
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

		// Correção sutil que estava no código original: extrair dados da resposta local
		registroNormalizado, err := conector.NormalizeData(corpoBytes, endpointConfig, nome)
		if err != nil {
			fmt.Printf("[Erro] Falha ao normalizar dados de %q: %v\n", nome, err)
			continue
		}
		registrosSequenciais[nome] = registroNormalizado
	}
	return registrosSequenciais
}

func main() {
	// -------------------------------------------------------------------------
	// ETAPA 1: Carregar Configurações YAML Dinamicamente (Infra de Conexão)
	// -------------------------------------------------------------------------
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("[Erro] Não foi possível encontrar o caminho do executável: %v\n", err)
		os.Exit(1)
	}

	execDir := filepath.Dir(execPath)
	apisPath := filepath.Join(execDir, "APIs")
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

	// Novo: Carrega os arquivos YAML que descrevem os Pipelines de Relatórios
	// Assume uma pasta chamada "Relatorios" ao lado do executável contendo as blueprints
	reportsPath := filepath.Join(execDir, "relatorios")
	reportsFS := os.DirFS(reportsPath)
	
	configRelatorios, err := gerador.LoadReportConfigFiles(reportsFS, "relatorio.yaml") // ou "." para varrer tudo
	if err != nil {
		fmt.Printf("[Erro] Falha ao carregar blueprints de relatórios: %v\n", err)
		os.Exit(1)
	}

	// -------------------------------------------------------------------------
	// ETAPA 2: Disparar Requisições Concorrentes usando URLs do YAML
	// -------------------------------------------------------------------------
	fmt.Println("\n[Debug] Iniciando Teste Comparativo de Coleta...")

	inicioSeq := time.Now()
	_ = carregarDadosSequencial(configEnv) 
	duracaoSeq := time.Since(inicioSeq)

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

	fmt.Printf("\n====== VEREDITO DE PERFORMANCE ======\n")
	fmt.Printf("Tempo no modo Sequencial : %v\n", duracaoSeq)
	fmt.Printf("Tempo no modo Concorrente: %v\n", duracaoPar)
	fmt.Printf("Ganho de Desempenho      : %.2f vezes mais rápido!\n", float64(duracaoSeq)/float64(duracaoPar))
	fmt.Printf("=====================================\n\n")

	// -------------------------------------------------------------------------
	// ETAPA 4: Correlacionar e Unificar os Dados baseados nas FKs (Correlacionador)
	// -------------------------------------------------------------------------
	fmt.Println("\n[Info] Iniciando a correlação de Chaves Estrangeiras (Joins)...")
	
	tabelaUnificada := correlacionador.MakeUnifiedData(todosOsRegistros, configEnv, rootTable)
	fmt.Printf("[Sucesso] Dados unificados com base na raiz %q. Total de linhas: %d\n", rootTable, len(tabelaUnificada))

	// -------------------------------------------------------------------------
	// ETAPA 5: Executar Pipelines e Exportar para Dashboard HTML Dinâmico
	// -------------------------------------------------------------------------
	fmt.Println("\n[Info] Iniciando processamento de métricas e geração do Dashboard HTML...")

	// Prepara a massa unificada bruta para a esteira funcional do Sumarizador
	dadosParaSumarizar := make([]sumarizador.JData, len(tabelaUnificada))
	for i, v := range tabelaUnificada {
		dadosParaSumarizar[i] = sumarizador.JData(v)
	}

	// Mapa que vai reter a saída final de cada relatório executado
	todosOsResultados := make(map[string][]sumarizador.JData)

	// Varre todos os relatórios declarados no YAML
	for nomeRelatorio, configReport := range configRelatorios {
		fmt.Printf("[Processando] Executando esteira para: %s\n", nomeRelatorio)
		
		// Executa as transformações sequenciais do pipeline passando o resultado do passo anterior
		resultadoPipeline := gerador.RunReportConfig(configReport, dadosParaSumarizar)
		
		// Registra o resultado final atrelado ao nome dele
		todosOsResultados[nomeRelatorio] = resultadoPipeline
	}

	// Caminho do arquivo de saída na raiz da aplicação
	htmlOutputPath := filepath.Join(execDir, "dashboard.html")

	// Dispara a geração da página rica em tabelas
	err = gerador.ExportToHTML(todosOsResultados, htmlOutputPath)
	if err != nil {
		fmt.Printf("[Erro] Falha catastrófica ao exportar relatório para HTML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n[OK] Pipeline executada com sucesso absoluto!\n")
	fmt.Printf("[Sucesso] Dashboard gerado com sucesso em: %s\n", htmlOutputPath)
	fmt.Printf("Tempo total decorrido: %v\n", time.Since(inicioSeq))
}
