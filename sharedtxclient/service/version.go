package service

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/shweini/dcrdtx/rpcserver/proto"
)

type VersionService struct {
    client pb.VersionServiceClient
}

func NewVersionService(conn *grpc.ClientConn) *VersionService {
    s := &VersionService{}
    s.client = pb.NewVersionServiceClient(conn)
    return s
}

func (v *VersionService) GetVersion() (string, error) {
    req := &pb.VersionRequest{}

    version, err := v.client.Version(context.Background(), req)
    if err != nil {
        return "", err
    }

    return version.String(), err
}
