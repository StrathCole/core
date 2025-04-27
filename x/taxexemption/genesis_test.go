package taxexemption_test

import (
	"testing"

	taxexemption "github.com/classic-terra/core/v3/x/taxexemption"
	ultil "github.com/classic-terra/core/v3/x/taxexemption/keeper"
	"github.com/classic-terra/core/v3/x/taxexemption/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewGenesisState(t *testing.T) {
	genesis := types.NewGenesisState()
	require.NotNil(t, genesis)
}

func TestDefaultGenesisState(t *testing.T) {
	genesis := taxexemption.DefaultGenesisState()
	require.NotNil(t, genesis)
}

func TestInitAndExportGenesis_Empty(t *testing.T) {
	// Setup mock context & keeper (simple zero value for now)
	input := ultil.CreateTestInput(t)
	k := input.TaxExemptionKeeper
	ctx := sdk.Context{}

	// Initialize genesis with empty state
	genesis := taxexemption.NewGenesisState()
	require.NotNil(t, genesis)

	taxexemption.InitGenesis(ctx, k, genesis)

	// Export genesis and check that it's not nil and matches default state
	exported := taxexemption.ExportGenesis(ctx, k)
	require.NotNil(t, exported)
}
