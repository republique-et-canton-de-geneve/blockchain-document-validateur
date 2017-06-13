package main

import (
	"flag"
	"log"
	"net"

	rcgech "github.com/Magicking/rc-ge-ch-pdf/internal/extract-services"
	"google.golang.org/grpc"
	//	"google.golang.org/grpc/grpclog"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:8090", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *serverAddr)
	if err != nil {
		log.Fatalf("fail to listen on %q: %v", serverAddr, err)
	}
	grpcServer := grpc.NewServer()

	rcgech.RegisterRegisterProxyServer(grpcServer, &rcgech.RcgechImplem{})

	grpcServer.Serve(lis)
}
