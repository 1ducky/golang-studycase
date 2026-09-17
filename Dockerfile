FROM golang:1.26

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .

ENV DB_USER=root
ENV DB_PASSWORD=secret
ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_NAME=todo_app

ENV SECRET_AUTH="ajdandjafjanalkdmakdalma dkakcsa"

ENV LOCAL_STORAGE_PATH="./storage/."

RUN go build -o server ./cmd/api/main.go

EXPOSE 8080

CMD ["./server"]