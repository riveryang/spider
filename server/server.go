// Copyright © 2015-2016 River Yang <comicme_yanghe@nanoframework.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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