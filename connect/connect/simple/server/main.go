package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	v1 "presentations/connect/pkg/simple/v1"
	"presentations/connect/pkg/simple/v1/simplev1connect"
	"time"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	names := []string{
		simplev1connect.SimpleServiceName,
	}
	reflector := grpcreflect.NewStaticReflector(names...)

	mux := http.NewServeMux()
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	path, handler := simplev1connect.NewSimpleServiceHandler(&Server{ideas: make(map[string]*v1.Idea)})
	mux.Handle(path, handler)

	log.Println("Starting Simple Connect Server")
	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		Addr:              "localhost:5555",
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

var _ simplev1connect.SimpleServiceHandler = &Server{}

type Server struct {
	ideas map[string]*v1.Idea
}

func (s Server) GetIdea(ctx context.Context, c *connect.Request[v1.GetIdeaRequest]) (*connect.Response[v1.Idea], error) {
	id, ok := s.ideas[c.Msg.GetName()]
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("could not find idea"))
	}
	return connect.NewResponse(id), nil
}

func (s Server) CreateIdea(ctx context.Context, c *connect.Request[v1.CreateIdeaRequest]) (*connect.Response[v1.Idea], error) {
	s.ideas[c.Msg.GetIdea().GetName()] = c.Msg.GetIdea()
	return connect.NewResponse(c.Msg.GetIdea()), nil
}

func (s Server) ListIdeas(ctx context.Context, c *connect.Request[v1.ListIdeasRequest]) (*connect.Response[v1.ListIdeasResponse], error) {
	ideas := make([]*v1.Idea, 0, len(s.ideas))
	for _, v := range s.ideas {
		ideas = append(ideas, v)
	}
	return connect.NewResponse(&v1.ListIdeasResponse{Ideas: ideas}), nil
}

func (s Server) DeleteIdea(ctx context.Context, c *connect.Request[v1.DeleteIdeaRequest]) (*connect.Response[v1.DeleteIdeaResponse], error) {
	delete(s.ideas, c.Msg.Name)
	return connect.NewResponse(&v1.DeleteIdeaResponse{}), nil
}

func (s Server) UpdateIdea(ctx context.Context, c *connect.Request[v1.UpdateIdeaRequest]) (*connect.Response[v1.Idea], error) {
	s.ideas[c.Msg.GetIdea().GetName()] = c.Msg.GetIdea()
	return connect.NewResponse(c.Msg.GetIdea()), nil
}
