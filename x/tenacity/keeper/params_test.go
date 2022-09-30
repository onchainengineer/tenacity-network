package keeper_test

import (
	"testing"

	testkeeper "github.com/onchainengineer/tenacity/testutil/keeper"
	"github.com/onchainengineer/tenacity/x/tenacity/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TenacityKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
