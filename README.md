
# Desafio Go: Stress Test

Este projeto em Go implementa um sistema para executar testes de carga em serviços web, permitindo configurar a URL alvo, o número total de requisições e o nível de concorrência desejado. Ao término dos testes, o sistema gera um relatório detalhado, que inclui informações sobre o tempo total gasto, o total de requisições realizadas e a distribuição dos status HTTP retornados.

## Pré-requisitos

Antes de começar, certifique-se de ter instalado em sua máquina:

- Go (versão 1.22.3 ou superior)
- Docker (opcional, para execução via contêiner Docker)

## Instalação e Configuração

### Clonando o repositório

```bash
git clone https://github.com/jonasjesusamerico/goexpert-stress-test.git
cd goexpert-stress-test
```

## Compilação e Execução

### Executando diretamente do código-fonte

1. **Compilação**:

   ```bash
   go build -o stresstester ./cmd
   ```

2. **Execução do teste de carga**:

   Substitua `--url`, `--requests` e `--concurrency` pelos valores adequados às suas necessidades.

   ```bash
   ./stresstester --url=http://google.com --requests=100 --concurrency=10
   ```

### Executando via Docker

1. **Construção da imagem Docker**:

   ```bash
   docker build -t goexpert-stress-test .
   ```

2. **Execução do teste de carga com Docker**:

   Substitua `--url`, `--requests` e `--concurrency` conforme necessário.

   ```bash
   docker run goexpert-stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

## Relatório de Resultados

Ao concluir o teste, o sistema apresenta um relatório detalhado com as seguintes informações:

- O tempo total das requisições
- A quantidde total de requisições realizadas
- A quantidade de requisições que foram bem-sucedidas (status HTTP 200)
- A exibição de outros códigos de status HTTP