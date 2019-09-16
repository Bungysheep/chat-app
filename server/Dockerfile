FROM golang

WORKDIR /app/chat-server

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o ./bin/chat-server ./cmd

EXPOSE 50051

CMD ["./bin/chat-server"]