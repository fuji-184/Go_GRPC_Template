package server_impls

import (
    "context"

    "google.golang.org/protobuf/types/known/emptypb"
    musik "github.com/fuji-184/streaming/musik"
)

type MusikServer struct {
    musik.UnimplementedMusikServer
}

func (s *MusikServer) InfoLagu(ctx context.Context, _ *emptypb.Empty) (*musik.Lagu, error) {
    return &musik.Lagu{
        Judul: "lagu fujiiii",
    }, nil
}


