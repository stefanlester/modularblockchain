package main

import (
	"fmt"
	"github.com/stefanlester/Modular-Blockchain/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	srv := network.NewServer(opts)
}

