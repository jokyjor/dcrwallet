package types

import (
    "github.com/decred/dcrd/wire"
    "github.com/decred/dcrd/chaincfg/chainhash"
    pb "github.com/shweini/dcrdtx/rpcserver/proto"
)

func TransactionResponse(res *pb.GetTransactionResponse) (*wire.MsgTx, error) {
    msgtx := wire.NewMsgTx()

    // Add All txouts
    for _,v := range res.Ticket.TxOut {
        msgtx.AddTxOut(wire.NewTxOut(v.Value, v.PkScript))
    }

    // Add all txins
    for _, v := range res.Ticket.TxIn {
        previousOp := v.PreviousOutPoint
        hash, err := chainhash.NewHashFromStr(previousOp.Hash)
        if err != nil {
            return nil, err
        }

        op := wire.NewOutPoint(hash, previousOp.Index, int8(previousOp.Tree))
        txin := wire.TxIn{
            PreviousOutPoint: *op,
            Sequence: v.Sequence,
            ValueIn: v.ValueIn,
            BlockHeight: v.BlockHeight,
            SignatureScript: v.SignatureScript,
        }
        msgtx.AddTxIn(&txin)
    }

    msgtx.Version = uint16(res.Ticket.Version)
    msgtx.SerType = wire.TxSerializeType(res.Ticket.SerType)
    msgtx.LockTime = res.Ticket.LockTime
    msgtx.Expiry = res.Ticket.Expiry

    return msgtx, nil
}


// TODO @jokyjor..young man, please abstract all similarities between this function and the one above
// into a single function and call as required
func TicketResponse(res *pb.PurchaseTicketResponse) (*wire.MsgTx, error) {
    msgtx := wire.NewMsgTx()

    // Add All txouts
    for _,v := range res.Ticket.TxOut {
        msgtx.AddTxOut(wire.NewTxOut(v.Value, v.PkScript))
    }

    // Add all txins
    for _, v := range res.Ticket.TxIn {
        previousOp := v.PreviousOutPoint
        hash, err := chainhash.NewHashFromStr(previousOp.Hash)
        if err != nil {
            return nil, err
        }

        op := wire.NewOutPoint(hash, previousOp.Index, int8(previousOp.Tree))
        txin := wire.TxIn{
            PreviousOutPoint: *op,
            Sequence: v.Sequence,
            ValueIn: v.ValueIn,
            BlockHeight: v.BlockHeight,
            SignatureScript: v.SignatureScript,
        }
        msgtx.AddTxIn(&txin)
    }

    msgtx.Version = uint16(res.Ticket.Version)
    msgtx.SerType = wire.TxSerializeType(res.Ticket.SerType)
    msgtx.LockTime = res.Ticket.LockTime
    msgtx.Expiry = res.Ticket.Expiry

    return msgtx, nil
}

func BuildTransactionRequest(hash *chainhash.Hash) (*pb.GetTransactionRequest, error) {
    return &pb.GetTransactionRequest{
        Hash: hash.String(),
    }, nil

}

func BuildPurchaseTicketRequest(msgtx *wire.MsgTx) (*pb.PurchaseTicketRequest, error) {
    req := &pb.PurchaseTicketRequest{
        Version: uint32(msgtx.Version),
        SerType: uint32(msgtx.SerType),
        LockTime: msgtx.LockTime,
        Expiry: msgtx.Expiry,
    }

    // Add all txins to the request
    for _, v := range msgtx.TxIn {
        op := &pb.OutPoint{
            Hash: v.PreviousOutPoint.Hash.String(),
            Index: v.PreviousOutPoint.Index,
            Tree: int32(v.PreviousOutPoint.Tree),
        }

        txin := &pb.TxIn{
            PreviousOutPoint: op,
            Sequence: v.Sequence,
            ValueIn: v.ValueIn,
            BlockHeight: v.BlockHeight,
            SignatureScript: v.SignatureScript,
        }
        req.TxIn = append(req.TxIn, txin)
    }

    // Add all txouts to the request
    for _, v := range msgtx.TxOut {
        txout := &pb.TxOut{
            Value: v.Value,
            Version: uint32(v.Version),
            PkScript: v.PkScript,
        }
        req.TxOut = append(req.TxOut, txout)
    }
    return req, nil
}



func getOutPoint(hash *chainhash.Hash, index uint32, tree uint32) *wire.OutPoint {
    return wire.NewOutPoint(hash, index, int8(tree))
}
