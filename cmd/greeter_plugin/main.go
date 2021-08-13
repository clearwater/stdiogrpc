//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/clearwater/stdiogrpc"
	"github.com/clearwater/stdiogrpc/cmd/helloworld"
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
		helloworld.RegisterGreeterServer(grpcServer, helloworld.NewServerImpl(log))
		reflection.Register(grpcServer)
		grpcServer.Serve(session)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gconn, err := grpc.Dial("stdio", grpc.WithInsecure(), grpc.WithContextDialer(session.Dial))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		grpcClient := helloworld.NewGreeterClient(gconn)
		// send messages
		for {
			helloworld.Greet(grpcClient, log, "host", "anne")
			time.Sleep(helloworld.Timeout)
		}
	}()

	wg.Wait()

}
