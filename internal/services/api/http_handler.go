package api

import (
	"encoding/json"
	"net/http"

	"github.com/yannis94/kanshi/internal/services/network"
)

type HTTPHandler struct {
	monitor *network.Monitor
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewHTTPHandler(m *network.Monitor) *HTTPHandler {
	return &HTTPHandler{monitor: m}
}

func (h HTTPHandler) GetBandwidth(w http.ResponseWriter, r *http.Request) {
	bytesByMilisecond, err := h.monitor.GetBandwidth("https://spin.atomicobject.com/wp-content/uploads/golang-logo.jpg")

	if err != nil {
		response := errorResponse{Code: 500, Message: err.Error()}
		body, _ := json.Marshal(response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(body)
		return
	}

	type Bandwidth struct {
		BytesByMilisecond int `json:"bytes_by_milisecond"`
	}

	w.Header().Add("Content-Type", "application/json")
	body, err := json.Marshal(APIResponse{Code: 200, Data: Bandwidth{BytesByMilisecond: bytesByMilisecond}})
	w.Write(body)
}

func (h HTTPHandler) GetNetworkInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := json.Marshal(APIResponse{Code: 200, Data: h.monitor.Network})
	w.Write(body)
}
