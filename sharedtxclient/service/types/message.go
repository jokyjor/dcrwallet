package types

import (
    "github.com/decred/dcrd/dcrutil"
    pb "github.com/shweini/dcrdtx/rpcserver/proto"
)

func BuildVerifyMessageRequest(address dcrutil.Address, message string, signature []byte) (*pb.VerifyMessageRequest, error) {
    req := &pb.VerifyMessageRequest{
        Address: address.String(),
        Message: message,
        Signature: signature,
    }

    return req, nil
}

func VerifyMessageResponse(res *pb.VerifyMessageResponse) (bool, error) {
    return res.Valid, nil
}
