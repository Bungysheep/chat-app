package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/bungysheep/chat-app/client/api/chat"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	hostname := "localhost"
	if os.Getenv("HOSTNAME") != "" {
		hostname = os.Getenv("HOSTNAME")
	}

	port := "50051"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	cliConn, err := grpc.Dial(fmt.Sprintf("%s:%s", hostname, port), opts...)
	if err != nil {
		log.Fatalf("Failed to dial with server: %v", err)
	}

	chatSvcCli := chat.NewChatServiceClient(cliConn)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	fmt.Printf("SendID:")
	sendID := ""
	if sc.Scan() {
		sendID = sc.Text()
	}

	fmt.Printf("SubID:")
	subID := ""
	if sc.Scan() {
		subID = sc.Text()
	}

	subCliDone := make(chan bool, 1)
	go func() {
		chatSubCli, err := chatSvcCli.Subscribe(ctx, &chat.SubscribeRequest{SubId: sendID})
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
			fmt.Printf("(%s): %s\n", resp.GetMessage().GetSendId(), resp.GetMessage().GetText())
		}
		subCliDone <- true
	}()

	interuptSignal := make(chan os.Signal, 1)
	signal.Notify(interuptSignal, os.Interrupt)

	for {
		select {
		case <-interuptSignal:
			unsubscribe(ctx, cliConn, chatSvcCli, sendID, subCliDone)
			return
		default:
			if sc.Scan() {
				text := sc.Text()
				if text != "" {
					switch text {
					case "<exit>":
						unsubscribe(ctx, cliConn, chatSvcCli, sendID, subCliDone)
						return
					default:
						_, err := chatSvcCli.Send(ctx, &chat.SendRequest{SubId: subID, Message: &chat.Message{SendId: sendID, Text: text}})
						if err != nil {
							log.Printf("Failed to send SendRequest: %v", err)
						}
					}
				}
			}
		}
	}
}

func unsubscribe(ctx context.Context, cliConn *grpc.ClientConn, chatSvcCli chat.ChatServiceClient, id string, subCliDone chan bool) {
	_, err := chatSvcCli.Send(ctx, &chat.SendRequest{SubId: id, Message: &chat.Message{Text: "<exit>"}})
	if err != nil {
		log.Printf("Failed to send SendRequest: %v", err)
	}

	<-subCliDone

	log.Printf("gRpc client is closing...\n")
	cliConn.Close()

	ctx.Done()
}
