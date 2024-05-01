package handlers

import (
	"encoding/json"
	"net/http"

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
	conn, err := grpc.Dial(":3002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return c.Render(http.StatusServiceUnavailable, "/views/error.html", map[string]interface{}{"code": http.StatusServiceUnavailable, "error": err.Error()})
	}

	defer conn.Close()

	client := pb.NewNetworkClient(conn)

	ntwInfo, err := client.GetNetworkInfo(c.Request().Context(), nil)

	if err != nil {
		return c.Render(http.StatusServiceUnavailable, "/views/error.html", map[string]interface{}{"code": http.StatusServiceUnavailable, "error": err.Error()})
	}

	var data models.Network

	if err := json.Unmarshal(ntwInfo.NetworkInfo, &data); err != nil {
		return c.Render(http.StatusInternalServerError, "/views/error.html", map[string]interface{}{"code": http.StatusInternalServerError, "error": err.Error()})
	}

	return c.Render(http.StatusOK, "/views/network.html", map[string]interface{}{"data": data})
}

func (h *NetworkHandler) GetBandwidth(c echo.Context) error {
	conn, err := grpc.Dial(":3002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return c.Render(http.StatusInternalServerError, "/views/error.html", map[string]interface{}{"code": http.StatusInternalServerError, "error": err.Error()})
	}

	defer conn.Close()

	client := pb.NewNetworkClient(conn)

	bandwidth, err := client.GetBandwidth(c.Request().Context(), nil)

	if err != nil {
		return c.Render(http.StatusInternalServerError, "/views/error.html", map[string]interface{}{"code": http.StatusInternalServerError, "error": err.Error()})
	}

	return c.Render(http.StatusOK, "/views/network.html", map[string]interface{}{"data": bandwidth.BytesPerMilisecond})
}

func (h *NetworkHandler) GetDevices(c echo.Context) error {
	conn, err := grpc.Dial(":3002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": http.StatusInternalServerError, "error": err.Error()})
	}

	defer conn.Close()

	client := pb.NewNetworkClient(conn)

	ntwInfo, err := client.GetNetworkInfo(c.Request().Context(), nil)

	if err != nil {
		return c.Render(http.StatusServiceUnavailable, "/views/error.html", map[string]interface{}{"code": http.StatusServiceUnavailable, "error": err.Error()})
	}

	var data models.Network

	if err := json.Unmarshal(ntwInfo.NetworkInfo, &data); err != nil {
		return c.Render(http.StatusInternalServerError, "/views/error.html", map[string]interface{}{"code": http.StatusInternalServerError, "error": err.Error()})
	}

	return c.HTMLBlob(http.StatusOK, data.Devices)
}
