# Análise Comparativa de Microsserviços: Go vs. Quarkus (JVM)

Este repositório contém um projeto de prova de conceito (PoC) para uma análise técnica e data-driven entre duas stacks de tecnologia para a construção de microsserviços: Go e Quarkus (rodando na JVM).

## 🎯 Objetivo

O objetivo principal deste projeto é avaliar e comparar as duas stacks com base nos seguintes critérios:
* **Experiência do Developer (DevEx):** Facilidade de configuração, velocidade do build local e ciclo de feedback (hot reload).
* **Performance:** Tempo de inicialização, consumo de memória e CPU em repouso e sob carga.
* **Artefactos de Build:** Tamanho das imagens de contêiner e complexidade do processo de deployment.

Os dados recolhidos servirão como base para decisões de arquitetura e para a introdução de novas tecnologias na equipa.

## 🛠️ Arquitetura

O projeto consiste em dois microsserviços funcionalmente idênticos, orquestrados via Docker Compose.

* **`go-service`**: Um web server simples escrito em Go puro, utilizando apenas a biblioteca padrão `net/http`.
* **`quarkus-service`**: Um web server implementado com Quarkus, utilizando as extensões `quarkus-rest` e `quarkus-rest-jackson`.

Ambos os serviços expõem um único endpoint de health check para fins de comparação.

## 📋 Pré-requisitos

Para executar este projeto, é necessário ter as seguintes ferramentas instaladas e a rodar na sua máquina:

* [Docker](https://www.docker.com/products/docker-desktop/)
* [Docker Compose](https://docs.docker.com/compose/) (geralmente já incluído no Docker Desktop)

Nenhuma outra dependência (Go, Java, Maven) precisa de ser instalada localmente.

## 🚀 Como Executar

Todo o ambiente é gerido pelo Docker Compose, garantindo consistência e uma configuração limpa.

1.  **Clone o Repositório:**
    ```bash
    git clone https://github.com/coelhodiana/go-vs-quarkus
    cd go-vs-quarkus
    ```

2.  **Construa e Inicie os Contêineres:**
    A partir da pasta raiz do projeto, execute o seguinte comando. Ele vai construir as imagens Docker e iniciar os contêineres. O tempo total do build será exibido no output do Docker.

    ```bash
    docker compose up --build
    ```

3.  **Para Parar os Serviços:**
    Pressione `Ctrl + C` no terminal onde os serviços estão a rodar. Para remover os contêineres, execute:
    ```bash
    docker compose down
    ```

## ✅ Verificação

Após a inicialização, os serviços estarão disponíveis nos seguintes endereços:

* **Go Service:**
    ```bash
    curl http://localhost:8081/health
    # Resposta esperada: Go service is up and running!
    ```

* **Quarkus Service:**
    ```bash
    curl http://localhost:8080/health
    # Resposta esperada: Quarkus service is up and running!
    ```

## 🧪 Testes e Qualidade

### `go-service`
* **Tecnologia:** Testes unitários com o pacote `testing` nativo do Go.
* **Execução Local:** Na pasta `go-service`, rode `make test`.
* **Relatório de Cobertura:** Na pasta `go-service`, rode `make coverage` para gerar um `coverage.html`.
* **Portão de Qualidade:** Os testes são executados automaticamente durante o `docker compose up --build`. Se um teste falhar, a imagem Docker do `go-app` não é criada.

### `quarkus-service`
* _(A ser implementado)_


## 📊 Análise

As observações e métricas recolhidas após a configuração inicial do projeto:

### Tempo de Build (Comando: `docker compose up --build`)
| Cenário | Tempo Total | Notas |
| :--- | :--- | :--- |
| **Build com Cache Frio** | **~83.3 segundos** | Simula a primeira execução num ambiente novo. A maior parte do tempo foi gasta a descarregar as dependências Maven do Quarkus. |
| **Build com Cache Quente (Sem Testes)** | **~1.1 segundos** | Simula um rebuild sem alterações no código, antes da adição dos testes ao build. |
| **Build com Cache Quente (Com Testes Go)** | **~2.0 segundos** | Mede o impacto de adicionar o portão de qualidade (`RUN go test`) ao build incremental. |

### Tempo de Inicialização da Aplicação (Startup)
* **Go:** Praticamente instantâneo (< 0.1s).
* **Quarkus (JVM):** Extremamente rápido para o padrão Java, registando um tempo de **~0.4s**.

### Tamanho da Imagem Final
A análise das imagens finais construídas pelo Docker revela uma diferença drástica, que impacta diretamente o tempo de deployment e o custo de armazenamento de artefactos.

| Serviço | Tamanho da Imagem | Notas |
| :--- | :--- | :--- |
| **`go-app`** | **~24 MB** | A imagem contém apenas o binário estático compilado e uma base mínima do Alpine Linux. Não há dependências externas. |
| **`quarkus-app`** | **~469 MB** | A imagem precisa de incluir a Java Runtime Environment (JRE) completa, além do JAR da aplicação e de todas as suas dependências. |


### Conclusões
* **Experiência do Developer (Build):** O ecossistema Go proporciona um ciclo de build drasticamente mais rápido, tanto com cache frio como quente. A complexidade do Maven e o número de dependências do Quarkus resultam num primeiro build significativamente mais lento.
* **Complexidade de Configuração:** O projeto Go exigiu menos ficheiros de configuração. O projeto Quarkus, embora gerado automaticamente, envolveu a resolução de mais problemas de configuração de build (versão do compilador, ficheiros ocultos, `.dockerignore`).
* **Tamanho da Imagem Docker:** A imagem do serviço Go (~24 MB) é quase **20 vezes mais leve** que a imagem do Quarkus/JVM (~469 MB), uma diferença crucial para a agilidade em pipelines de CI/CD e custos de armazenamento.
