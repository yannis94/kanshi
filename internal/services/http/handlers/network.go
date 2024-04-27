package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/yannis94/kanshi/internal/models"
	pb "github.com/yannis94/kanshi/internal/proto/network_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NetworkHandler struct{}

func NewNetworkHandler() *NetworkHandler {
	return &NetworkHandler{}
}

func (h *NetworkHandler) GetInfo(c echo.Context) error {
	conn, err := grpc.Dial("localhost:3002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	defer conn.Close()

	client := pb.NewNetworkClient(conn)

	ntwInfo, err := client.GetNetworkInfo(c.Request().Context(), nil)

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	var data models.Network

	if err := json.Unmarshal(ntwInfo.NetworkInfo, &data); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, data)
}

func (h *NetworkHandler) GetBandwidth(c echo.Context) error {
	conn, err := grpc.Dial("localhost:3002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	defer conn.Close()

	client := pb.NewNetworkClient(conn)

	bandwidth, err := client.GetBandwidth(c.Request().Context(), nil)

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]int32{"bandwidth": bandwidth.BytesPerMilisecond})
}
