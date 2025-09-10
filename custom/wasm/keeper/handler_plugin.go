package keeper

import (
	treasurykeeper "github.com/classic-terra/core/v3/x/treasury/keeper"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	taxexemptionkeeper "github.com/classic-terra/core/v3/x/taxexemption/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	taxkeeper "github.com/classic-terra/core/v3/x/tax/keeper"
	taxtypes "github.com/classic-terra/core/v3/x/tax/types"
)

// msgEncoder is an extension point to customize encodings
type msgEncoder interface {
	// Encode converts wasmvm message to n cosmos message types
	Encode(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Msg, error)
}

// MessageRouter ADR 031 request type routing
type MessageRouter interface {
	Handler(msg sdk.Msg) baseapp.MsgServiceHandler
}

// SDKMessageHandler can handles messages that can be encoded into sdk.Message types and routed.
type SDKMessageHandler struct {
	router             MessageRouter
	encoders           msgEncoder
	treasuryKeeper     treasurykeeper.Keeper
	accountKeeper      authkeeper.AccountKeeper
	bankKeeper         bankKeeper.Keeper
	taxexemptionKeeper taxexemptionkeeper.Keeper
	taxKeeper          taxkeeper.Keeper
}

func NewMessageHandler(
	router MessageRouter,
	ics4Wrapper wasmtypes.ICS4Wrapper,
	channelKeeper wasmtypes.ChannelKeeper,
	capabilityKeeper wasmtypes.CapabilityKeeper,
	bankKeeper bankKeeper.Keeper,
	taxexemptionKeeper taxexemptionkeeper.Keeper,
	treasuryKeeper treasurykeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	taxKeeper taxkeeper.Keeper,
	unpacker codectypes.AnyUnpacker,
	portSource wasmtypes.ICS20TransferPortSource,
	customEncoders ...*wasmkeeper.MessageEncoders,
) wasmkeeper.Messenger {
	encoders := wasmkeeper.DefaultEncoders(unpacker, portSource)
	for _, e := range customEncoders {
		encoders = encoders.Merge(e)
	}
	return wasmkeeper.NewMessageHandlerChain(
		NewSDKMessageHandler(router, encoders, taxexemptionKeeper, treasuryKeeper, accountKeeper, bankKeeper, taxKeeper),
		wasmkeeper.NewBurnCoinMessageHandler(bankKeeper),
	)
}

func NewSDKMessageHandler(router MessageRouter, encoders msgEncoder, taxexemptionKeeper taxexemptionkeeper.Keeper, treasuryKeeper treasurykeeper.Keeper, accountKeeper authkeeper.AccountKeeper, bankKeeper bankKeeper.Keeper, taxKeeper taxkeeper.Keeper) SDKMessageHandler {
	return SDKMessageHandler{
		router:             router,
		encoders:           encoders,
		treasuryKeeper:     treasuryKeeper,
		taxexemptionKeeper: taxexemptionKeeper,
		accountKeeper:      accountKeeper,
		bankKeeper:         bankKeeper,
		taxKeeper:          taxKeeper,
	}
}

func (h SDKMessageHandler) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) (events []sdk.Event, data [][]byte, msgs [][]*codectypes.Any, err error) {
	sdkMsgs, err := h.encoders.Encode(ctx, contractAddr, contractIBCPortID, msg)
	if err != nil {
		return nil, nil, nil, err
	}

	// contract handling is ALWAYS reverse charged
	ctx = ctx.WithValue(taxtypes.ContextKeyTaxReverseCharge, true)

	for _, sdkMsg := range sdkMsgs {
		// Charge tax on result msg
		// we set simulate to false here as it is not available and we don't need to
		// increase the tax amount for simulation inside of wasm
		res, err := h.handleSdkMessage(ctx, contractAddr, sdkMsg)
		if err != nil {
			return nil, nil, nil, err
		}
		// append data
		data = append(data, res.Data)
		// append events
		sdkEvents := make([]sdk.Event, len(res.Events))
		for i := range res.Events {
			sdkEvents[i] = sdk.Event(res.Events[i])
		}
		events = append(events, sdkEvents...)
		// no additional msg responses to return from SDK handler
	}
	return events, data, nil, nil
}

func (h SDKMessageHandler) handleSdkMessage(ctx sdk.Context, contractAddr sdk.Address, msg sdk.Msg) (*sdk.Result, error) {
	// Route and execute the message; message basic validation and signer checks are handled by the router endpoints
	if handler := h.router.Handler(msg); handler != nil {
		msgResult, err := handler(ctx, msg)
		return msgResult, err
	}
	return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "can't route message %+v", msg)
}
