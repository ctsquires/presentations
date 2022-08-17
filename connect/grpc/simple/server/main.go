package main

import (
	"context"
	"log"
	"net"
	v1 "presentations/connect/pkg/simple/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	grpcServer := grpc.NewServer()

	v1.RegisterSimpleServiceServer(grpcServer, &Server{ideas: make(map[string]*v1.Idea)})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "localhost:4444")
	if err != nil {
		panic(err)
	}

	log.Println("Starting gRPC Simple Server on localhost:4444")

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

var _ v1.SimpleServiceServer = &Server{}

type Server struct {
	v1.UnimplementedSimpleServiceServer
	ideas map[string]*v1.Idea
}

func (s *Server) GetIdea(ctx context.Context, request *v1.GetIdeaRequest) (*v1.Idea, error) {
	id, ok := s.ideas[request.GetName()]
	if !ok {
		return nil, status.Error(codes.NotFound, "could not find idea")
	}
	return id, nil
}

func (s *Server) CreateIdea(ctx context.Context, request *v1.CreateIdeaRequest) (*v1.Idea, error) {
	s.ideas[request.GetIdea().GetName()] = request.GetIdea()
	return request.GetIdea(), nil
}

func (s *Server) ListIdeas(ctx context.Context, request *v1.ListIdeasRequest) (*v1.ListIdeasResponse, error) {
	ideas := make([]*v1.Idea, 0, len(s.ideas))
	for _, v := range s.ideas {
		ideas = append(ideas, v)
	}
	return &v1.ListIdeasResponse{Ideas: ideas}, nil
}

func (s *Server) DeleteIdea(ctx context.Context, request *v1.DeleteIdeaRequest) (*v1.DeleteIdeaResponse, error) {
	delete(s.ideas, request.Name)
	return &v1.DeleteIdeaResponse{}, nil
}

func (s *Server) UpdateIdea(ctx context.Context, request *v1.UpdateIdeaRequest) (*v1.Idea, error) {
	s.ideas[request.GetIdea().GetName()] = request.GetIdea()
	return request.GetIdea(), nil
}
