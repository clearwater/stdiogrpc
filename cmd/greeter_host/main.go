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
	"github.com/clearwater/stdiogrpc/cmd/hostproto"
	"github.com/clearwater/stdiogrpc/cmd/pluginproto"
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
		hostproto.RegisterHostServer(grpcServer, hostproto.NewServerImpl(log))
		reflection.Register(grpcServer)
		grpcServer.Serve(session)
	}()

	wg.Add(1)
	go func() {
		const timeout = 2 * time.Second
		// create and run client
		defer wg.Done()
		gconn, err := grpc.Dial("stdio", grpc.WithInsecure(), grpc.WithContextDialer(session.Dial))
		if err != nil {
			log.Fatalln("failed to create grpc client: ", err)
		}
		defer gconn.Close()

		grpcClient := pluginproto.NewPluginClient(gconn)
		for {
			// send messages
			pluginproto.CallPlugin(grpcClient, log, "plugin", "Bob")
			time.Sleep(timeout)
		}
	}()

	wg.Wait()
}
