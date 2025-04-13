package client_impls

import (
    "context"
    "log"

    "google.golang.org/protobuf/types/known/emptypb"
    musik_pb "github.com/fuji-184/streaming/musik"
)

func RunMusik(ctx context.Context, musik_client musik_pb.MusikClient) error {
    lagu, err := musik_client.InfoLagu(ctx, &emptypb.Empty{})
    if err != nil {
        log.Fatalf("failed to get lagu: %v", err)
    }

    log.Printf("lagu: %s", lagu.GetJudul())

    return nil
}
