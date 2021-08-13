//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/clearwater/stdiogrpc"
	"github.com/clearwater/stdiogrpc/cmd/helloworld"
)

func main() {
	log := log.New(os.Stderr, "host   ", log.Ltime)
	log.Printf("Host Starting\n")

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s plugincmd ...args", os.Args[0])
	}

	// plugin exec and args are in os.Args[1:]
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	session, err := stdiogrpc.NewHostSession(cmd)
	if err != nil {
		panic(err)
	}
	cmd.Stderr = os.Stderr
	cmd.Start()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// create and run server
		defer wg.Done()
		grpcServer := grpc.NewServer()
		helloworld.RegisterGreeterServer(grpcServer, helloworld.NewServerImpl(log))
		reflection.Register(grpcServer)
		grpcServer.Serve(session)
	}()

	wg.Add(1)
	go func() {
		// create and run client
		defer wg.Done()
		gconn, err := grpc.Dial("stdio", grpc.WithInsecure(), grpc.WithContextDialer(session.Dial))
		if err != nil {
			log.Fatalln("failed to create grpc client: ", err)
		}
		defer gconn.Close()

		grpcClient := helloworld.NewGreeterClient(gconn)
		for {
			// send messages
			helloworld.Greet(grpcClient, log, "plugin", "bob")
			time.Sleep(helloworld.Timeout)
		}
	}()

	wg.Wait()
}
