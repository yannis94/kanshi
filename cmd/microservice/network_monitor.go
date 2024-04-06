package main

import (
	"fmt"

	"github.com/yannis94/kanshi/internal/services/network"
	"github.com/yannis94/kanshi/internal/store"
)

func main() {
	fmt.Println("####### KANSHI #######")
	fmt.Println("Network monioring")
	nms := store.NewNetworkMemoryStore()

	m := network.NewMonitor(nms, "home")

	fmt.Printf("%+v\n", m)
}
