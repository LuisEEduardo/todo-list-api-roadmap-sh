# Definindo a imagem base
FROM golang:1.23

# Criando diretório dentro da imagem que está sendo construída
WORKDIR /app  

# Copiando os arquivos de dependências do Go para o
# diretório de trabalho dentro da imagem
COPY go.mod go.sum ./

# Baixando as dependências do Go
RUN go mod download

# Copiando o restante dos arquivos da aplicação para o
# diretório de trabalho dentro da imagem
COPY . .

# Compilando a aplicação Go
RUN go build -v -o /usr/local/bin/app ./cmd/api

# Definindo o comando padrão para executar a aplicação
CMD ["app"]