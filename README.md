# stdiogrpc
## Bidirectional gRPC with a Subprocess Using stdio

stdiogrpc is a library supporting bidirectional gRPC between a host process and a subprocess over stdio.
It uses hashicorp/yamux to support multiple connections over a single channel.

Examples are included showing how to communicate with a plugin loaded a subprocess.

The following extracts from the included samples show the key concepts.

## Create a session on Host-side
```go
// create a new process
cmd := exec.Command(cmdLine[0], cmdLine[1:]...)

// create a new session binding stdin+stdout from the subprocess
session, err := stdiogrpc.NewHostSession(cmd)
if err != nil {
	panic(err)
}

// map stderr from the child process to my stderr
cmd.Stderr = os.Stderr

// start the child process
err = cmd.Start()
if err != nil {
	panic(err)
}
```

## Create a session Plugin-side
```go
// create a new session binding stdin+stdout
session, err := stdiogrpc.NewPluginSession()
if err != nil {
	panic(err)
}
```

## Create a gRPC server on host and plugin
```go
grpcServer := grpc.NewServer()
hostproto.RegisterHostServer(grpcServer, hostproto.NewServerImpl(log))
reflection.Register(grpcServer)
go grpcServer.Serve(session)  // pass the stdiogrpc.Session here
```

## Make gRPC calls from host or plugin
```go
gconn, err := grpc.Dial("stdio", grpc.WithInsecure(), grpc.WithContextDialer(session.Dial))
if err != nil {
	log.Fatalf("did not connect: %v", err)
}
grpcClient := hostproto.NewHostClient(gconn)

// send messages
hostproto.CallHost(grpcClient, log, "host", "Anne")
```