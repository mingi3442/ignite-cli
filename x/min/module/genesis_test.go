package min_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "min/testutil/keeper"
	"min/testutil/nullify"
	"min/x/min/module"
	"min/x/min/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MinKeeper(t)
	min.InitGenesis(ctx, k, genesisState)
	got := min.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
