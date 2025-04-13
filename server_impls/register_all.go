package server_impls

import (
    "google.golang.org/grpc"

    hello "github.com/fuji-184/streaming/proto"
    musik "github.com/fuji-184/streaming/musik"
)

func RegisterAllServices(s *grpc.Server){
    hello.RegisterHelloServer(s, &HelloServer{})
    musik.RegisterMusikServer(s, &MusikServer{})
}
