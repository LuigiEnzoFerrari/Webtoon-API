FROM golang:1.18

WORKDIR app/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
ARG PORT_B=8080
ENV PORT=$PORT_B
EXPOSE $PORT_B
COPY . .
RUN go mod init mymod
RUN go mod tidy

ENTRYPOINT go run main.go
