package server_impls

import (
    "log"
    "context"

    pb "github.com/fuji-184/streaming/proto"
)

type HelloServer struct {
    pb.UnimplementedHelloServer
}

func (s *HelloServer) Menyapa(ctx context.Context, req *pb.Isi) (*pb.Hasil, error) {
    log.Printf("Received: %s", req.GetPesan())
    return &pb.Hasil{
        Balasan: "hallo jugaaa",
    }, nil
}

func (s *HelloServer) MenyapaStreaming(req *pb.Isi, stream pb.Hello_MenyapaStreamingServer) error {
    for i := 0; i<5; i++ {
        stream.Send(&pb.Hasil{
            Balasan: "hello jugaaa",
        })
    }
    return nil
}

