package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bungysheep/chat-app/client/api/chat"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	cliConn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to dial with server: %v", err)
	}

	chatSvcCli := chat.NewChatServiceClient(cliConn)

	subCliDone := make(chan bool, 1)
	go func() {
		chatSubCli, err := chatSvcCli.Subscribe(ctx, &chat.SubscribeRequest{})
		if err != nil {
			log.Printf("Failed to stream SubscribeResponse: %v", err)
		}

		for {
			resp, err := chatSubCli.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("Failed to receive SubscribeResponse: %v", err)
				break
			}
			fmt.Printf("(SERVER): %s\n", resp.GetMessage().GetText())
		}
		subCliDone <- true
	}()

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for {
		select {
		default:
			if sc.Scan() {
				text := sc.Text()
				if text != "" {
					switch text {
					case "<exit>":
						_, err := chatSvcCli.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: "<exit>"}})
						if err != nil {
							log.Printf("Failed to send SendRequest: %v", err)
						}

						<-subCliDone

						log.Printf("gRpc client is closing...\n")
						cliConn.Close()

						ctx.Done()
						return
					default:
						_, err := chatSvcCli.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: text}})
						if err != nil {
							log.Printf("Failed to send SendRequest: %v", err)
						}
					}
				}
			}
		}
	}
}
