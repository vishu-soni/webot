package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	botpb "webot/proto/bot/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var channelName = flag.String("usercode", "default", "Channel name for chatting")
var senderType = flag.String("usertype", "default", "Senders name")
var tcpServer = flag.String("server", ":9005", "Tcp server")

func joinChannel(ctx context.Context, client botpb.ChatServiceClient) {

	channel := botpb.Channel{
		UserCode: *channelName,
		UserType: *senderType,
		Status:   true,
	}
	stream, err := client.JoinChannel(ctx, &channel)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}

	fmt.Printf("Joined channel: %v \n", *channelName)

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
			}

			if *channelName != in.Sender {
				fmt.Printf("MESSAGE: (%v) -> %v \n", in.Sender, in.Message)
			}
		}
	}()

}

func sendMessage(ctx context.Context, client botpb.ChatServiceClient, message string) {
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := botpb.Message{
		Sender: *channelName,
		Channel: &botpb.Channel{
			UserCode: *channelName,
			UserType: *senderType,
			Status:   true,
		},
		Message: message,
	}
	stream.Send(&msg)

	ack, err := stream.CloseAndRecv()
	if err != nil {
		stream.Context().Done()
	}
	fmt.Printf("Message sent: %v \n", ack)
}

func main() {

	flag.Parse()

	fmt.Println("--- CLIENT APP ---")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := botpb.NewChatServiceClient(conn)

	go joinChannel(ctx, client)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, scanner.Text())
	}

}
