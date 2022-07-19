package keeper

import (
	"github.com/ecostation-labs/ecostation-chain/x/ecostationchain/types"
)

var _ types.QueryServer = Keeper{}
