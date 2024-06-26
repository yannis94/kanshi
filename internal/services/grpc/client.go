package grpc

import (
	"context"
	"encoding/json"

	pb "github.com/yannis94/kanshi/internal/proto/network_grpc"
	"github.com/yannis94/kanshi/internal/services/network"
)

type NetworkGRPCServer struct {
	Monitor *network.Monitor
}

// GetBandwidth implements network_grpc.NetworkServer.
func (s NetworkGRPCServer) GetBandwidth(context.Context, *pb.GetBandwidthRequest) (*pb.GetBandwidthResponse, error) {
	bytesByMilisecond, err := s.Monitor.GetBandwidth("https://spin.atomicobject.com/wp-content/uploads/golang-logo.jpg")

	return &pb.GetBandwidthResponse{BytesPerMilisecond: int32(bytesByMilisecond)}, err
}

// GetNetworkInfo implements network_grpc.NetworkServer.
func (s NetworkGRPCServer) GetNetworkInfo(context.Context, *pb.GetNetworkInfoRequest) (*pb.GetNetworkInfoResponse, error) {
	data, err := json.Marshal(s.Monitor.Network)

	if err != nil {
		return nil, err
	}
	return &pb.GetNetworkInfoResponse{NetworkInfo: data}, nil
}

// mustEmbedUnimplementedNetworkServer implements network_grpc.NetworkServer.
func (s NetworkGRPCServer) mustEmbedUnimplementedNetworkServer() {
	panic("unimplemented")
}
