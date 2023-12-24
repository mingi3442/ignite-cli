package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "min/testutil/keeper"
	"min/x/min/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MinKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
