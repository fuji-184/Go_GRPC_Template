package client_impls

import (
    "context"
    "log"
    "google.golang.org/grpc"

    hello_pb "github.com/fuji-184/streaming/proto"
    musik_pb "github.com/fuji-184/streaming/musik"
)

func RunAllClients(conn *grpc.ClientConn){

    hello := hello_pb.NewHelloClient(conn)
    musik := musik_pb.NewMusikClient(conn)

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    if err := RunHello(ctx, hello); err != nil {
        log.Fatalf("failed to run hello service: %v", err)
    }

    if err := RunMusik(ctx, musik); err != nil {
        log.Fatalf("failed to run musik service: %v", err)
    }

}
