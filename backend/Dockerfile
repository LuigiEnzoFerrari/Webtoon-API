FROM golang:1.18

WORKDIR app/

ARG PORT_B=8080
ENV PORT=$PORT_B
EXPOSE $PORT_B
COPY . .
RUN go mod init mymod
RUN go mod tidy

ENTRYPOINT go run main.go
