package server

import (
	"golang.org/x/net/context"
	"github.com/riveryang/spider/pb"
	"time"
	"log"
)

type Server struct {
	clients map[string] time.Time
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if s.clients == nil {
		s.clients = make(map[string] time.Time)
	}

	s.clients[req.Address] = time.Now()
	log.Printf("Register or heartbeat node: %v", req.Address)
	return &pb.RegisterResponse{Message: "OK"}, nil
}

func (s *Server) Ttl() {
	for address, t := range s.clients {
		if time.Now().Unix() - t.Unix() > 10 {
			log.Printf("Deleted expired node: %v", address)
			delete(s.clients, address)
		}
	}
}