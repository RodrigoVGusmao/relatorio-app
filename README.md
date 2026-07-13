# 📊 Go Pipeline de Relatórios Analíticos (Log Engine)

Este projeto é um motor de processamento analítico de dados extremamente rápido escrito em **Go**. Ele realiza a coleta concorrente de dados a partir de APIs simuladas (mocks), executa normalização dinâmica com base em definições YAML, realiza a correlação relacional em memória (Joins usando FKs) e entrega relatórios agregados e consolidados em um **Dashboard HTML interativo**.

---

## 🏗️ Arquitetura do Projeto

O sistema foi modularizado de forma robusta utilizando os seguintes pacotes:

```
.
├── APIs/                  # Configurações de conexão e mapeamento de dados (YAML)
├── conector/              # Motor de descoberta, consumo concorrente e normalização
├── correlacionador/       # Lógica de correlação em grafo (Joins utilizando FKs)
├── gerador/               # Parser de relatórios YAML e motor de renderização HTML
├── mockAPIs/              # Massa de dados de teste (JSON) de usuários, redes e logs
├── relatorios/            # Blueprints (YAML) que definem os pipelines de agregação
├── sumarizador/           # Funções puras de estatística (Sum, Count, Avg, Min, Max)
├── Dockerfile             # Build otimizada em multi-stage para Alpine Linux
└── docker-compose.yml     # Orquestração do app e dos 4 servidores mock locais
```

---

## ⚡ Diferenciais Técnicos

* **Coleta Concorrente:** O módulo `conector` consome múltiplos endpoints em paralelo utilizando `sync.WaitGroup` e `channels`, obtendo um ganho de performance drástico frente à abordagem sequencial.
* **Segurança de Tipos:** Conversão de schemas dinâmicos baseados em `map[string]any`, resolvendo problemas de mutabilidade de memória nativa do Go através de uma arquitetura funcional de cópias rasas (*shallow copy*).
* **Filtros Robustos e Estáticos:** Estrutura estéril que impede conflitos sintáticos ao renderizar strings complexas contendo caminhos com pontos (ex: `logs.porta_destino`).

---

## 🛠️ Como Executar o Projeto (Via Docker Compose)

A infraestrutura inteira roda encapsulada pelo Docker, garantindo que as dependências de rede e ambiente estejam isoladas.

### Requisitos
* `Docker` e `Docker Compose` instalados.

### Passos para Inicialização

1. **Crie o arquivo de destino local para sincronismo de arquivo:**
   ```bash
   touch dashboard.html
   ```

2. **Inicie os serviços via docker-compose:**
   ```bash
   docker compose up --build app
   ```

3. **O que acontece ao rodar:**
   * O Compose sobe 4 servidores mock simulando APIs REST locais populadas com dados relacionais sintéticos.
   * O container `app` aguarda a saúde dos mocks ficarem `healthy`.
   * Ele inicia e faz a requisição concorrente para todas as APIs.
   * Exibe no terminal um **Veredito de Desempenho** comparando a velocidade concorrente contra a sequencial.
   * O correlacionador une as tabelas em uma massa única baseada no nó pai raiz.
   * O processador executa os pipelines declarados em `relatorios/relatorio.yaml`.
   * Grava em tempo real os resultados no arquivo `./dashboard.html` mapeado diretamente em sua máquina local.

---

## 🧪 Executando os Testes Unitários

Caso queira validar os submódulos e garantir que nenhum bug de agregação ou normalização passou despercebido:

```bash
go test ./...
```

---

## 📊 Visualizando os Resultados

Assim que o container finalizar a execução, você encontrará um arquivo chamado `dashboard.html` na raiz do seu projeto. 

Basta abrir este arquivo no seu navegador favorito para acompanhar o relatório formatado, responsivo e limpo, contendo blocos visuais automáticos das seguintes operações:
1. **Acessos por criticidade**
2. **Payload total de MB por porta de rede**
3. **Maior payload total trafegado**
