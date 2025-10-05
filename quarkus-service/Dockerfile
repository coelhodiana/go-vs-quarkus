# ESTÁGIO 1: Build da Aplicação (O Construtor com Maven)
# Usamos uma imagem que contém o Maven e o Java JDK para compilar o projeto
FROM maven:3.9-eclipse-temurin-21 AS build

# Define o nosso diretório de trabalho
WORKDIR /app

# Copiamos primeiro os ficheiros do Maven para descarregar as dependências.
# Isto aproveita o sistema de cache do Docker de forma eficiente.
COPY .mvn/ .mvn/
COPY mvnw pom.xml ./
RUN ./mvnw dependency:go-offline

# Agora, copiamos o código-fonte da nossa aplicação
COPY src ./src

# Finalmente, compilamos o projeto. O resultado será um JAR executável.
# A opção -DskipTests acelera o build por não executar os testes.
RUN ./mvnw package -DskipTests


# ESTÁGIO 2: Imagem Final (O Corredor com JRE)
# Começamos de uma imagem Java leve, apenas com o necessário para executar (JRE).
FROM eclipse-temurin:21-jre

# Define o diretório de trabalho final
WORKDIR /deploy

# Expõe a porta padrão do Quarkus
EXPOSE 8080

# Copia a aplicação já compilada (o JAR e as suas dependências) do estágio de build
COPY --from=build /app/target/quarkus-app ./

# Comando final para iniciar a aplicação Quarkus
CMD ["java", "-jar", "quarkus-run.jar"]