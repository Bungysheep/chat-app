FROM golang

WORKDIR /app/chat-client

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./bin/chat-client ./cmd

CMD ["./bin/chat-client"]