package tenacity_test

import (
	"testing"

	keepertest "github.com/onchainengineer/tenacity/testutil/keeper"
	"github.com/onchainengineer/tenacity/testutil/nullify"
	"github.com/onchainengineer/tenacity/x/tenacity"
	"github.com/onchainengineer/tenacity/x/tenacity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TenacityKeeper(t)
	tenacity.InitGenesis(ctx, *k, genesisState)
	got := tenacity.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
