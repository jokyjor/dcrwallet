package service

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"

    "github.com/decred/dcrd/wire"
    "github.com/decred/dcrd/chaincfg/chainhash"

    pb "github.com/shweini/dcrdtx/rpcserver/proto"
    "github.com/decred/dcrwallet/sharedtxclient/service/types"
)

type TransactionService struct {
    client pb.TransactionServiceClient
}

func NewTransactionService(conn *grpc.ClientConn) *TransactionService {
    t := &TransactionService{}
    t.client = pb.NewTransactionServiceClient(conn)

    return t
}

func (t *TransactionService) GetTransaction(hash *chainhash.Hash) (*wire.MsgTx, error) {
    req, err := types.BuildTransactionRequest(hash)
    if err != nil {
        return nil, err
    }

    tx, err := t.client.GetTransaction(context.Background(), req)
    if err != nil {
        return nil, err
    }

    msgtx, err := types.TransactionResponse(tx)
    return msgtx, err
}

func (t *TransactionService) PurchaseTicket(msgtx *wire.MsgTx) (*wire.MsgTx, error) {
    req, err := types.BuildPurchaseTicketRequest(msgtx)
    if err != nil {
        return nil, err
    }

    tx, err := t.client.PurchaseTicket(context.Background(), req)
    if err != nil {
        return nil, err
    }

    ticket, err := types.TicketResponse(tx)
    return ticket, nil
}
