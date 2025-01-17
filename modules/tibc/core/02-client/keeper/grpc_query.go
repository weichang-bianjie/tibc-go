package keeper

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	host "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	"github.com/bianjieai/tibc-go/modules/tibc/core/exported"
)

var _ types.QueryServer = Keeper{}

// ClientState implements the Query/ClientState gRPC method
func (q Keeper) ClientState(c context.Context, req *types.QueryClientStateRequest) (*types.QueryClientStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := host.ClientIdentifierValidator(req.ChainName); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(c)
	clientState, found := q.GetClientState(ctx, req.ChainName)
	if !found {
		return nil, status.Error(
			codes.NotFound,
			sdkerrors.Wrap(types.ErrClientNotFound, req.ChainName).Error(),
		)
	}

	any, err := types.PackClientState(clientState)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	proofHeight := types.GetSelfHeight(ctx)
	return &types.QueryClientStateResponse{
		ClientState: any,
		ProofHeight: proofHeight,
	}, nil
}

// ClientStates implements the Query/ClientStates gRPC method
func (q Keeper) ClientStates(c context.Context, req *types.QueryClientStatesRequest) (*types.QueryClientStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	clientStates := types.IdentifiedClientStates{}
	store := prefix.NewStore(ctx.KVStore(q.storeKey), host.KeyClientStorePrefix)

	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		keySplit := strings.Split(string(key), "/")
		if keySplit[len(keySplit)-1] != "clientState" {
			return nil
		}

		clientState, err := q.UnmarshalClientState(value)
		if err != nil {
			return err
		}

		chainName := keySplit[1]
		if err := host.ClientIdentifierValidator(chainName); err != nil {
			return err
		}

		identifiedClient := types.NewIdentifiedClientState(chainName, clientState)
		clientStates = append(clientStates, identifiedClient)
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Sort(clientStates)

	return &types.QueryClientStatesResponse{
		ClientStates: clientStates,
		Pagination:   pageRes,
	}, nil
}

// ConsensusState implements the Query/ConsensusState gRPC method
func (q Keeper) ConsensusState(c context.Context, req *types.QueryConsensusStateRequest) (*types.QueryConsensusStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := host.ClientIdentifierValidator(req.ChainName); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(c)

	var (
		consensusState exported.ConsensusState
		found          bool
	)

	height := types.NewHeight(req.RevisionNumber, req.RevisionHeight)
	if req.LatestHeight {
		consensusState, found = q.GetLatestClientConsensusState(ctx, req.ChainName)
	} else {
		if req.RevisionHeight == 0 {
			return nil, status.Error(codes.InvalidArgument, "consensus state height cannot be 0")
		}

		consensusState, found = q.GetClientConsensusState(ctx, req.ChainName, height)
	}

	if !found {
		return nil, status.Error(
			codes.NotFound,
			sdkerrors.Wrapf(types.ErrConsensusStateNotFound, "chain-name: %s, height: %s", req.ChainName, height).Error(),
		)
	}

	any, err := types.PackConsensusState(consensusState)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	proofHeight := types.GetSelfHeight(ctx)
	return &types.QueryConsensusStateResponse{
		ConsensusState: any,
		ProofHeight:    proofHeight,
	}, nil
}

// ConsensusStates implements the Query/ConsensusStates gRPC method
func (q Keeper) ConsensusStates(c context.Context, req *types.QueryConsensusStatesRequest) (*types.QueryConsensusStatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := host.ClientIdentifierValidator(req.ChainName); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(c)

	consensusStates := []types.ConsensusStateWithHeight{}
	store := prefix.NewStore(ctx.KVStore(q.storeKey), host.FullClientKey(req.ChainName, []byte(fmt.Sprintf("%s/", host.KeyConsensusStatePrefix))))

	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		// filter any metadata stored under consensus state key
		if strings.Contains(string(key), "/") {
			return false, nil
		}

		height, err := types.ParseHeight(string(key))
		if err != nil {
			return false, err
		}

		consensusState, err := q.UnmarshalConsensusState(value)
		if err != nil {
			return false, err
		}

		consensusStates = append(consensusStates, types.NewConsensusStateWithHeight(height, consensusState))
		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryConsensusStatesResponse{
		ConsensusStates: consensusStates,
		Pagination:      pageRes,
	}, nil
}

// Relayers implements the Query/Relayers gRPC method
func (q Keeper) Relayers(c context.Context, req *types.QueryRelayersRequest) (*types.QueryRelayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if err := host.ClientIdentifierValidator(req.ChainName); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryRelayersResponse{
		Relayers: q.GetRelayers(ctx, req.ChainName),
	}, nil
}
