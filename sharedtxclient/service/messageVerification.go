package service

import (
    "golang.org/x/net/context"
    "github.com/decred/dcrwallet/sharedtxclient/service/types"

    "github.com/decred/dcrd/dcrutil"
    "google.golang.org/grpc"
    pb "github.com/shweini/dcrdtx/rpcserver/proto"
)

type MessageVerificationService struct {
    client pb.MessageVerificationServiceClient
}

// NewMessageVerificationService
func NewMessageVerificationService(conn *grpc.ClientConn) *MessageVerificationService {
    s := &MessageVerificationService{}
    s.client = pb.NewMessageVerificationServiceClient(conn)
    return s
}

func (m *MessageVerificationService) VerifyMessage(address dcrutil.Address, message string, signature []byte) (bool, error) {
    req, err := types.BuildVerifyMessageRequest(address, message, signature)
    if err != nil {
        return false, err
    }

    res, err := m.client.VerifyMessage(context.Background(), req)
    if err != nil {
        return false, err
    }

    return types.VerifyMessageResponse(res)
}
