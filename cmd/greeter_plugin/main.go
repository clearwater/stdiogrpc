//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/clearwater/stdiogrpc"
	"github.com/clearwater/stdiogrpc/cmd/hostproto"
	"github.com/clearwater/stdiogrpc/cmd/pluginproto"
)

func main() {

	log := log.New(os.Stderr, "plugin ", log.Ltime)
	log.Printf("Client Starting\n")

	session, err := stdiogrpc.NewPluginSession()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// create server
		grpcServer := grpc.NewServer()
		pluginproto.RegisterPluginServer(grpcServer, pluginproto.NewServerImpl(log))
		grpcServer.Serve(session)
	}()

	wg.Add(1)
	go func() {
		const timeout = 1 * time.Second
		defer wg.Done()
		gconn, err := grpc.Dial("stdio", grpc.WithInsecure(), grpc.WithContextDialer(session.Dial))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		grpcClient := hostproto.NewHostClient(gconn)
		// send messages
		for {
			hostproto.CallHost(grpcClient, log, "host", "Anne")
			time.Sleep(timeout)
		}
	}()

	wg.Wait()

}
