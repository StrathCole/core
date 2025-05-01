package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdklegacy "github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/classic-terra/core/v3/x/taxexemption/types/legacy"
)

// RegisterInterfaces associates protoName with the new message types
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgAddTaxExemptionZone{},
		&MsgRemoveTaxExemptionZone{},
		&MsgModifyTaxExemptionZone{},
		&MsgAddTaxExemptionAddress{},
		&MsgRemoveTaxExemptionAddress{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&legacy.AddTaxExemptionZoneProposal{},
		&legacy.RemoveTaxExemptionZoneProposal{},
		&legacy.ModifyTaxExemptionZoneProposal{},
		&legacy.AddTaxExemptionAddressProposal{},
		&legacy.RemoveTaxExemptionAddressProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers the concrete types on the Amino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdklegacy.RegisterAminoMsg(cdc, &MsgAddTaxExemptionZone{}, "taxexemption/AddTaxExemptionZone")
	sdklegacy.RegisterAminoMsg(cdc, &MsgRemoveTaxExemptionZone{}, "taxexemption/RemoveTaxExemptionZone")
	sdklegacy.RegisterAminoMsg(cdc, &MsgModifyTaxExemptionZone{}, "taxexemption/ModifyTaxExemptionZone")
	sdklegacy.RegisterAminoMsg(cdc, &MsgAddTaxExemptionAddress{}, "taxexemption/AddTaxExemptionAddress")
	sdklegacy.RegisterAminoMsg(cdc, &MsgRemoveTaxExemptionAddress{}, "taxexemption/RemoveTaxExemptionAddress")
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
