FROM golang

WORKDIR /app/chat-server

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./bin/chat-server ./cmd

CMD ["./bin/chat-server"]