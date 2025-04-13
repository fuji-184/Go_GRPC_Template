package client_impls

import (
    "context"
    "log"
    "io"

    hello_pb "github.com/fuji-184/streaming/proto"
)

func RunHello(ctx context.Context, hello_client hello_pb.HelloClient) error {

    stream, err := hello_client.MenyapaStreaming(ctx, &hello_pb.Isi{
        Pesan: "helloooo",
    })
    if err != nil {
        log.Fatalf("failed to get Menyapa: %v", err)
    }

    for {
        hasil, err := stream.Recv()
        if err == io.EOF {
            // stream selesai
            break
        }
        if err != nil {
            log.Fatalf("failed to get stream data: %v", err)
        }

        log.Printf("hasil: %s", hasil.GetBalasan())
    }

    return nil
}
