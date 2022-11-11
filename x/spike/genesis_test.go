package spike_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "spike/testutil/keeper"
	"spike/testutil/nullify"
	"spike/x/spike"
	"spike/x/spike/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SpikeKeeper(t)
	spike.InitGenesis(ctx, *k, genesisState)
	got := spike.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
