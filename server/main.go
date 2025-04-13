package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "github.com/fuji-184/streaming/server_impls"
)

func main(){
    creds, err := credentials.NewServerTLSFromFile("tls/server.crt", "tls/server.key")
    if err != nil {
        log.Fatalf("failed to load tls files for server: %v", err)
    }

    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer(
        grpc.Creds(creds),
        grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
            log.Printf("unary interceptor: %s", info.FullMethod)
            return handler(ctx, req)
        }),
    )

    server_impls.RegisterAllServices(s)

    log.Printf("server started on :50051")

    if err := s.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
