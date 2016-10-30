// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"github.com/riveryang/spider/pb"
	"golang.org/x/net/context"
	"log"
	"net"
	"strconv"
	"time"
)

// nodeCmdCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := cmd.Flags().GetString("server")
		if err != nil {
			return err
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetUint("port")
		if err != nil {
			return err
		}

		address, listen := bind(host, int(port))
		connect, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer connect.Close()

		registry := pb.NewRegistryClient(connect)
		go func (registry pb.RegistryClient) {
			for ; ;  {
				res, err := registry.Register(context.Background(), &pb.RegisterRequest{Address: address})
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("%v", res.Message)
				time.Sleep(time.Second * 5)
			}
		}(registry)

		s := grpc.NewServer()
		log.Printf("Starting node with: %v", address)
		if err := s.Serve(listen); err != nil {
			log.Fatal("Start node err")
			return err
		}

		return nil
	},
}

func bind(host string, port int) (string, net.Listener) {
	port += 1
	address := host + ":" + strconv.Itoa(port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return bind(host, port)
	}

	return address, listen
}

func init() {
	RootCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("server", "s", "127.0.0.1", "Spider server address")
	nodeCmd.Flags().String("host", "127.0.0.1", "Spider node host")
	nodeCmd.Flags().UintP("port", "p", 6160, "Spider node port")
}
