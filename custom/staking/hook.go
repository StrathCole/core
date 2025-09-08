package staking

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	ColumbusChainID = "columbus-5"
)

var _ stakingtypes.StakingHooks = &TerraStakingHooks{}

// TerraStakingHooks implements staking hooks to enforce validator power limit
type TerraStakingHooks struct {
	sk stakingkeeper.Keeper
}

func NewTerraStakingHooks(sk stakingkeeper.Keeper) *TerraStakingHooks {
	return &TerraStakingHooks{sk: sk}
}

// Implement required staking hooks interface methods
func (h TerraStakingHooks) BeforeDelegationCreated(_ context.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h TerraStakingHooks) BeforeDelegationSharesModified(_ context.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// Other required hook methods with empty implementations
func (h TerraStakingHooks) AfterDelegationModified(ctx context.Context, _ sdk.AccAddress, valAddr sdk.ValAddress) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ChainID() != ColumbusChainID {
		return nil
	}

	validator, err := h.sk.GetValidator(ctx, valAddr)
	if err != nil {
		return nil
	}

	// Get validator's current power (after delegation modified)
	validatorPower := sdk.TokensToConsensusPower(validator.Tokens, h.sk.PowerReduction(ctx))

	// Get the total power of the validator set
	totalPower, _ := h.sk.GetLastTotalPower(ctx)
	if totalPower.IsZero() {
		return nil
	}

	// Get validator delegation percent
	validatorDelegationPercent := math.LegacyNewDec(validatorPower).QuoInt64(totalPower.Int64())

	if validatorDelegationPercent.GT(math.LegacyNewDecWithPrec(20, 2)) {
		return fmt.Errorf("validator power is over the allowed limit")
	}

	return nil
}

func (h TerraStakingHooks) BeforeValidatorSlashed(_ context.Context, _ sdk.ValAddress, _ math.LegacyDec) error {
	return nil
}

func (h TerraStakingHooks) BeforeValidatorModified(_ context.Context, _ sdk.ValAddress) error {
	return nil
}

func (h TerraStakingHooks) AfterValidatorBonded(_ context.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h TerraStakingHooks) AfterValidatorBeginUnbonding(_ context.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h TerraStakingHooks) AfterValidatorRemoved(_ context.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h TerraStakingHooks) AfterUnbondingInitiated(_ context.Context, _ uint64) error {
	return nil
}

// Add this method to TerraStakingHooks
func (h TerraStakingHooks) AfterValidatorCreated(_ context.Context, _ sdk.ValAddress) error {
	return nil
}

// Add the missing method
func (h TerraStakingHooks) BeforeDelegationRemoved(_ context.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}
