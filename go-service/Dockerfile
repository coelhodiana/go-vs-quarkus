# ESTÁGIO 1: Build da Aplicação (O Construtor)
# Usamos uma imagem com todas as ferramentas do Go
FROM golang:1.22-alpine AS builder

# Define o nosso diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiamos os ficheiros de módulo primeiro para aproveitar o cache do Docker.
# Se as dependências não mudarem, o Docker não as descarrega novamente.
COPY go.mod go.sum* ./
RUN go mod download

# Agora copiamos o resto do código-fonte
COPY *.go ./

# Executa os testes. Se falharem, o build para aqui.
RUN go test ./...

# Compilamos a nossa aplicação num único ficheiro binário estático.
# Isto é otimizado para correr em Linux (o ambiente padrão dos contêineres).
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app-binary .


# ESTÁGIO 2: Imagem Final (O Corredor)
# Começamos de uma imagem super pequena, apenas com o essencial.
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /app

# Copiamos APENAS o ficheiro binário que compilámos no estágio anterior.
COPY --from=builder /go-app-binary .

# "Avisamos" ao Docker que a nossa aplicação vai usar a porta 8081.
EXPOSE 8081

# O comando para executar a nossa aplicação quando o contêiner iniciar.
CMD ["./go-app-binary"]