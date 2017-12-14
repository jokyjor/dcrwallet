package types

import (
    "github.com/decred/dcrd/wire"
    //"github.com/decred/dcrd/dcrutil"
    "github.com/decred/dcrwallet/wallet/udb"
    pb "github.com/shweini/dcrdtx/rpcserver/proto"
)

func SignCreditsResponse(res *pb.SignCreditsResponse) ([]udb.Credit, error) {
    return nil, nil
}

func BuildSignCreditsRequest(credit []udb.Credit, ticket *wire.MsgTx) (*pb.SignCreditsRequest, error) {
    req := &pb.SignCreditsRequest{}

    for _,v := range credit {
        c := &pb.Credits{
            OutPoint: getRequestOutPoint(v.OutPoint),
            Amount: int64(v.Amount),
            PkScript: v.PkScript,
            Received: v.Received.UnixNano(),
            FromCoinBase: v.FromCoinBase,
        }
        req.Credits = append(req.Credits, c)
    }


    return req, nil
}

func getRequestOutPoint(op wire.OutPoint) *pb.OutPoint {
    return &pb.OutPoint{
        Hash: op.Hash.String(),
        Index: op.Index,
        Tree: int32(op.Tree),
    }
}
