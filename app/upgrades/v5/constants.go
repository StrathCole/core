package v5

import (
	store "cosmossdk.io/store/types"
	"github.com/classic-terra/core/v3/app/upgrades"

	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
)

const UpgradeName = "v5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV5UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			ibcfeetypes.StoreKey,
			icacontrollertypes.StoreKey,
		},
	},
}
