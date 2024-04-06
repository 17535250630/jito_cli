package geyser_client

import (
	"context"
	"os"
	"testing"

	"github.com/17535250630/jito_cli/proto"
	"github.com/stretchr/testify/assert"
)

func Test_GeyserClient(t *testing.T) {
	rpcAddr, ok := os.LookupEnv("GEYSER_RPC")
	assert.True(t, ok, "getting JITO_RPC from .env")

	client, err := NewGeyserClient(
		context.Background(),
		rpcAddr,
		nil,
	)
	assert.NoError(t, err)

	ctx := context.Background()

	t.Run("SubscribeBlockUpdates", func(t *testing.T) {
		sub, err := client.SubscribeBlockUpdates()
		assert.NoError(t, err)

		ch := make(chan *proto.BlockUpdate)
		client.OnBlockUpdates(ctx, sub, ch)

		block := <-ch
		assert.NotNil(t, block.BlockHeight)
	})
}
