package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/yannis94/kanshi/internal/proto/network_grpc"
	"github.com/yannis94/kanshi/internal/services/api"
	kanshi_grpc "github.com/yannis94/kanshi/internal/services/grpc"
	"github.com/yannis94/kanshi/internal/services/network"
	"github.com/yannis94/kanshi/internal/store"
	"google.golang.org/grpc"
)

func main() {
	var (
		GRPCServer, RESTServer bool
		port                   int
	)

	flag.BoolVar(&GRPCServer, "grpc", true, "gRPC server address")
	flag.BoolVar(&RESTServer, "rest", false, "REST server address")
	flag.IntVar(&port, "p", 3002, "Port to listen on")

	nms := store.NewNetworkMemoryStore()

	fmt.Println("####### KANSHI #######")
	fmt.Println("Network monitoring")

	if GRPCServer {
		listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		monitor := network.NewMonitor(nms, "home")

		grpcServer := grpc.NewServer()
		pb.RegisterNetworkServer(grpcServer, &kanshi_grpc.NetworkGRPCServer{Monitor: monitor})

		fmt.Println("gRPC server starting on port", port)
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}
	if RESTServer {
		monitor := network.NewMonitor(nms, "home")
		handler := api.NewHTTPHandler(monitor)
		http.HandleFunc("GET /bandwidth", handler.GetBandwidth)
		http.HandleFunc("GET /network", handler.GetNetworkInfo)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}

	fmt.Println("Quit")
	fmt.Println("####### KANSHI #######")
}
