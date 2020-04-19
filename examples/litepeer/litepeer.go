package main

// This example launches an IPFS-Lite peer and fetches a hello-world
// hash from the IPFS network.

import (
	"bytes"
	"context"
	"fmt"

	sdkc "github.com/RTradeLtd/go-temporalx-sdk/client"

	ipfslite "github.com/RTradeLtd/ipfs-lite"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Bootstrappers are using 1024 keys. See:
	// https://github.com/ipfs/infra/issues/378
	crypto.MinRsaKeyBits = 1024

	ds, err := ipfslite.BadgerDatastore("test")
	if err != nil {
		panic(err)
	}
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		panic(err)
	}

	listen, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")

	h, dht, err := ipfslite.SetupLibp2p(
		ctx,
		priv,
		nil,
		[]multiaddr.Multiaddr{listen},
		ds,
		ipfslite.Libp2pOptionsExtra...,
	)

	if err != nil {
		panic(err)
	}

	client, err := sdkc.NewClient(sdkc.Opts{
		ListenAddress: "xapi.temporal.cloud:9090",
		Insecure:      true,
	})
	if err != nil {
		panic(err)
	}
	lite, err := ipfslite.New(ctx, zap.NewNop(), ds, h, dht, nil, client)
	if err != nil {
		panic(err)
	}

	//	lite.Bootstrap(ipfslite.DefaultBootstrapPeers())
	nd, err := lite.AddFile(ctx, bytes.NewReader([]byte("helo world")), nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", nd)
}
