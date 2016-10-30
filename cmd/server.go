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

package cmd

import (
	"github.com/spf13/cobra"
	"net"
	"strconv"
	"google.golang.org/grpc"
	"github.com/riveryang/spider/pb"
	"github.com/riveryang/spider/server"
	"time"
	"log"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

		port, err := cmd.Flags().GetUint("port")
		if err != nil {
			return err
		}

		address := host + ":" + strconv.Itoa(int(port))
		listen, err := net.Listen("tcp", address)
		if err != nil {
			return err
		}

		s := grpc.NewServer()
		spiderServer := &server.Server{}
		pb.RegisterRegistryServer(s, spiderServer)
		go func (spiderServer *server.Server) {
			for ; ;  {
				spiderServer.Ttl()
				time.Sleep(time.Second)
			}
		}(spiderServer)

		log.Printf("Starting server with: %v", address)
		if err := s.Serve(listen); err != nil {
			log.Fatal("Start server err")
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.Flags().String("host", "127.0.0.1", "Spider server host")
	serverCmd.Flags().UintP("port", "p", 6060, "Spider server port")
}
