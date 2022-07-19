package ecostationchain_test

import (
	"testing"

	keepertest "github.com/ecostation-labs/ecostation-chain/testutil/keeper"
	"github.com/ecostation-labs/ecostation-chain/testutil/nullify"
	"github.com/ecostation-labs/ecostation-chain/x/ecostationchain"
	"github.com/ecostation-labs/ecostation-chain/x/ecostationchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EcostationchainKeeper(t)
	ecostationchain.InitGenesis(ctx, *k, genesisState)
	got := ecostationchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
