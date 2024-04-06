package pkg

import (
	"github.com/17535250630/jito_cli/proto"
	"github.com/17535250630/solana-go"
	bin "github.com/gagliardetto/binary"
)

// ConvertTransactionToProtobufPacket converts a solana-go Transaction to a proto.Packet.
func ConvertTransactionToProtobufPacket(transaction *solana.Transaction) (proto.Packet, error) {
	data, err := transaction.MarshalBinary()
	if err != nil {
		return proto.Packet{}, err
	}

	return proto.Packet{
		Data: data,
		Meta: &proto.Meta{
			Size:        uint64(len(data)),
			Addr:        "",
			Port:        0,
			Flags:       nil,
			SenderStake: 0,
		},
	}, nil
}

// ConvertBatchTransactionToProtobufPacket converts a slice of solana-go Transaction to a slice of proto.Packet.
func ConvertBatchTransactionToProtobufPacket(transactions []*solana.Transaction) ([]*proto.Packet, error) {
	packets := make([]*proto.Packet, 0, len(transactions))
	for _, tx := range transactions {
		packet, err := ConvertTransactionToProtobufPacket(tx)
		if err != nil {
			return nil, err
		}

		packets = append(packets, &packet)
	}

	return packets, nil
}

// ConvertProtobufPacketToTransaction converts a proto.Packet to a solana-go Transaction.
func ConvertProtobufPacketToTransaction(packet *proto.Packet) (*solana.Transaction, error) {
	tx := &solana.Transaction{}
	err := tx.UnmarshalWithDecoder(bin.NewBorshDecoder(packet.Data))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// ConvertBatchProtobufPacketToTransaction converts a slice of proto.Packet to a slice of solana-go Transaction.
func ConvertBatchProtobufPacketToTransaction(packets []*proto.Packet) ([]*solana.Transaction, error) {
	txs := make([]*solana.Transaction, 0, len(packets))
	for _, packet := range packets {
		tx, err := ConvertProtobufPacketToTransaction(packet)
		if err != nil {
			return nil, err
		}

		txs = append(txs, tx)
	}

	return txs, nil
}
