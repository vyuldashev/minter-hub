package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.NewLegacyAmino()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterInterfaces regiesteres the interfaces for the proto stuff
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgValsetConfirm{},
		&MsgSendToEth{},
		&MsgRequestBatch{},
		&MsgConfirmBatch{},
		&MsgDepositClaim{},
		&MsgWithdrawClaim{},
		&MsgSendToMinterClaim{},
		&MsgSetOrchestratorAddress{},
	)

	registry.RegisterInterface(
		"peggy.v1beta1.EthereumClaim",
		(*EthereumClaim)(nil),
		&MsgDepositClaim{},
		&MsgWithdrawClaim{},
		&MsgSendToMinterClaim{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*EthereumClaim)(nil), nil)
	cdc.RegisterConcrete(&MsgSetOrchestratorAddress{}, "peggy/MsgSetOrchestratorAddress", nil)
	cdc.RegisterConcrete(&MsgValsetConfirm{}, "peggy/MsgValsetConfirm", nil)
	cdc.RegisterConcrete(&MsgSendToEth{}, "peggy/MsgSendToEth", nil)
	cdc.RegisterConcrete(&MsgRequestBatch{}, "peggy/MsgRequestBatch", nil)
	cdc.RegisterConcrete(&MsgConfirmBatch{}, "peggy/MsgConfirmBatch", nil)
	cdc.RegisterConcrete(&Valset{}, "peggy/Valset", nil)
	cdc.RegisterConcrete(&MsgDepositClaim{}, "peggy/MsgDepositClaim", nil)
	cdc.RegisterConcrete(&MsgWithdrawClaim{}, "peggy/MsgWithdrawClaim", nil)
	cdc.RegisterConcrete(&MsgSendToMinterClaim{}, "peggy/MsgSendToMinterClaim", nil)
	cdc.RegisterConcrete(&OutgoingTxBatch{}, "peggy/OutgoingTxBatch", nil)
	cdc.RegisterConcrete(&OutgoingTransferTx{}, "peggy/OutgoingTransferTx", nil)
	cdc.RegisterConcrete(&ERC20Token{}, "peggy/ERC20Token", nil)
	cdc.RegisterConcrete(&IDSet{}, "peggy/IDSet", nil)
	cdc.RegisterConcrete(&Attestation{}, "peggy/Attestation", nil)
}
