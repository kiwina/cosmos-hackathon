package goldcdp

import (
	"github.com/bandprotocol/goldcdp/x/goldcdp/keeper"
	"github.com/bandprotocol/goldcdp/x/goldcdp/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier

	NewProduct          = types.NewProduct
	NewMsgCreateProduct = types.NewMsgCreateProduct
	NewMsgUpdateProduct = types.NewMsgUpdateProduct

	NewMsgCreateSell = types.NewMsgCreateSell
	NewMsgUpdateSell = types.NewMsgUpdateSell
	NewMsgDeleteSell = types.NewMsgDeleteSell

	NewMsgCreateReservation = types.NewMsgCreateReservation
	NewMsgUpdateReservation = types.NewMsgUpdateReservation
	NewMsgDeleteReservation = types.NewMsgDeleteReservation
)

type (
	Keeper              = keeper.Keeper
	MsgBuyGold          = types.MsgBuyGold
	MsgSetSourceChannel = types.MsgSetSourceChannel

	Product          = types.Product
	MsgCreateProduct = types.MsgCreateProduct
	MsgUpdateProduct = types.MsgUpdateProduct

	Sell          = types.Sell
	MsgCreateSell = types.MsgCreateSell
	MsgUpdateSell = types.MsgUpdateSell
	MsgDeleteSell = types.MsgDeleteSell

	Reservation          = types.Reservation
	MsgCreateReservation = types.MsgCreateReservation
	MsgUpdateReservation = types.MsgUpdateReservation
	MsgDeleteReservation = types.MsgDeleteReservation
)
