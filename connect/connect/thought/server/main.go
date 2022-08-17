package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	v1 "presentations/connect/pkg/thought/v1"
	"presentations/connect/pkg/thought/v1/thoughtv1connect"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	connect_go "github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	names := []string{
		thoughtv1connect.ThoughtServiceName,
	}
	reflector := grpcreflect.NewStaticReflector(names...)
	checker := grpchealth.NewStaticChecker(names...)

	mux := http.NewServeMux()
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	mux.Handle(grpchealth.NewHandler(checker))

	path, handler := thoughtv1connect.NewThoughtServiceHandler(&Server{})
	mux.Handle(path, handler)

	log.Println("Starting Connect Server")
	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		Addr:              "localhost:6789",
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

var _ thoughtv1connect.ThoughtServiceHandler = &Server{}

type Server struct {
}

func (s *Server) GetThought(ctx context.Context, c *connect_go.Request[v1.GetThoughtRequest]) (*connect_go.Response[v1.GetThoughtResponse], error) {
	return connect_go.NewResponse(&v1.GetThoughtResponse{Thought: gofakeit.HipsterSentence(gofakeit.IntRange(5, 10))}), nil
}

func (s *Server) ReceiveThoughts(ctx context.Context, c *connect_go.Request[v1.ReceiveThoughtsRequest], c2 *connect_go.ServerStream[v1.ReceiveThoughtsResponse]) error {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			if err := c2.Send(&v1.ReceiveThoughtsResponse{Thought: gofakeit.HipsterSentence(gofakeit.IntRange(5, 10))}); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (s *Server) SendThoughts(ctx context.Context, c *connect_go.ClientStream[v1.SendThoughtsRequest]) (*connect_go.Response[v1.SendThoughtsResponse], error) {
	// TODO implement me
	panic("implement me")
}

func (s *Server) ConverseThoughts(ctx context.Context, c *connect_go.BidiStream[v1.ConverseThoughtsRequest, v1.ConverseThoughtsResponse]) error {
	for {
		req, err := c.Receive()
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
		if err := c.Send(&v1.ConverseThoughtsResponse{
			Thought: fmt.Sprintf("What do I think of %s? Well %s", req.GetThought(), gofakeit.HipsterSentence(gofakeit.IntRange(5, 15))),
		}); err != nil {
			return err
		}
	}

}
