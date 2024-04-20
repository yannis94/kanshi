package main

import (
	"fmt"
	"net/http"

	"github.com/yannis94/kanshi/internal/services/api"
	"github.com/yannis94/kanshi/internal/services/network"
	"github.com/yannis94/kanshi/internal/store"
)

func main() {
	fmt.Println("####### KANSHI #######")
	fmt.Println("Network monioring")
	nms := store.NewNetworkMemoryStore()

	monitor := network.NewMonitor(nms, "home")
	handler := api.NewHTTPHandler(monitor)
	http.HandleFunc("GET /bandwidth", handler.GetBandwidth)
	http.HandleFunc("GET /network", handler.GetNetworkInfo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
