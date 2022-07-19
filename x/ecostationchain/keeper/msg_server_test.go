package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/ecostation-labs/ecostation-chain/testutil/keeper"
	"github.com/ecostation-labs/ecostation-chain/x/ecostationchain/keeper"
	"github.com/ecostation-labs/ecostation-chain/x/ecostationchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EcostationchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
