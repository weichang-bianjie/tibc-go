package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/simulation"
	host "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	"github.com/bianjieai/tibc-go/simapp"
)

func TestDecodeStore(t *testing.T) {
	app := simapp.Setup(false)
	cdc := app.AppCodec()

	channelID := "channelidone"
	portID := "portidone"

	bz := []byte{0x1, 0x2, 0x3}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   host.NextSequenceSendKey(portID, channelID),
				Value: sdk.Uint64ToBigEndian(1),
			},
			{
				Key:   host.PacketCommitmentKey(portID, channelID, 1),
				Value: bz,
			},
			{
				Key:   host.PacketAcknowledgementKey(portID, channelID, 1),
				Value: bz,
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"NextSeqSend", "NextSeqSend A: 1\nNextSeqSend B: 1"},
		{"NextSeqRecv", "NextSeqRecv A: 1\nNextSeqRecv B: 1"},
		{"NextSeqAck", "NextSeqAck A: 1\nNextSeqAck B: 1"},
		{"CommitmentHash", fmt.Sprintf("CommitmentHash A: %X\nCommitmentHash B: %X", bz, bz)},
		{"AckHash", fmt.Sprintf("AckHash A: %X\nAckHash B: %X", bz, bz)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			res, found := simulation.NewDecodeStore(cdc, kvPairs.Pairs[i], kvPairs.Pairs[i])
			if i == len(tests)-1 {
				require.False(t, found, string(kvPairs.Pairs[i].Key))
				require.Empty(t, res, string(kvPairs.Pairs[i].Key))
			} else {
				require.True(t, found, string(kvPairs.Pairs[i].Key))
				require.Equal(t, tt.expectedLog, res, string(kvPairs.Pairs[i].Key))
			}
		})
	}
}
