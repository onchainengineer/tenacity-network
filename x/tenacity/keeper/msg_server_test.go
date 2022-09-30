package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/onchainengineer/tenacity/testutil/keeper"
	"github.com/onchainengineer/tenacity/x/tenacity/keeper"
	"github.com/onchainengineer/tenacity/x/tenacity/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TenacityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
