package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/bianjieai/tibc-go/modules/tibc/core/02-client/simulation"
	"github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	host "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	ibctmtypes "github.com/bianjieai/tibc-go/modules/tibc/light-clients/07-tendermint/types"
	"github.com/bianjieai/tibc-go/simapp"
)

func TestDecodeStore(t *testing.T) {
	app := simapp.Setup(false)
	chainName := "clientidone"

	height := types.NewHeight(0, 10)

	clientState := &ibctmtypes.ClientState{}
	consState := &ibctmtypes.ConsensusState{
		Timestamp: time.Now().UTC(),
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   host.FullClientStateKey(chainName),
				Value: app.TIBCKeeper.ClientKeeper.MustMarshalClientState(clientState),
			},
			{
				Key:   host.FullConsensusStateKey(chainName, height),
				Value: app.TIBCKeeper.ClientKeeper.MustMarshalConsensusState(consState),
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
		{"ClientState", fmt.Sprintf("ClientState A: %v\nClientState B: %v", clientState, clientState)},
		{"ConsensusState", fmt.Sprintf("ConsensusState A: %v\nConsensusState B: %v", consState, consState)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			res, found := simulation.NewDecodeStore(app.TIBCKeeper.ClientKeeper, kvPairs.Pairs[i], kvPairs.Pairs[i])
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
