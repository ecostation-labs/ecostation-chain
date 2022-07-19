package keeper_test

import (
	"testing"

	testkeeper "github.com/ecostation-labs/ecostation-chain/testutil/keeper"
	"github.com/ecostation-labs/ecostation-chain/x/ecostationchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EcostationchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
