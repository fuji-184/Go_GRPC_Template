package main

import (
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"

    "github.com/fuji-184/streaming/client_impls"
)

func main(){
    creds, err := credentials.NewClientTLSFromFile("tls/server.crt", "")
    if err != nil {
        log.Fatalf("failed to load tls files for client: %v", err)
    }

    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()

    client_impls.RunAllClients(conn)
}
