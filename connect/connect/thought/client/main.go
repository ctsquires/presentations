package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	v1 "presentations/connect/pkg/thought/v1"
	"presentations/connect/pkg/thought/v1/thoughtv1connect"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	connect_go "github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
)

func main() {
	client := thoughtv1connect.NewThoughtServiceClient(&http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}, "http://localhost:6789")

	fmt.Println("Getting a Random Thought")
	time.Sleep(1 * time.Second)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := client.GetThought(ctx, connect_go.NewRequest(&v1.GetThoughtRequest{}))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received Thought: %s\n", res.Msg.Thought)

	fmt.Println("Going to converse thoughts")

	stream := client.ConverseThoughts(context.Background())
	go func() {
		for {
			res, err := stream.Receive()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res.Thought)
		}
	}()

	for start := time.Now(); time.Since(start) < 20*time.Second; {
		thought := gofakeit.NounAbstract()
		fmt.Printf("What about %s\n", thought)
		if err := stream.Send(&v1.ConverseThoughtsRequest{Thought: thought}); err != nil && err != io.EOF {
			panic(err)
		}
		time.Sleep(4 * time.Second)
	}

	fmt.Println("Thanks for the chat, closing stream")

	if err := stream.CloseRequest(); err != nil {
		panic(err)
	}
}
