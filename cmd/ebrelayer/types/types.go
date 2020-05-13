package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Event enum containing supported chain events
type Event byte

const (
	// Unsupported is an invalid Cosmos or Ethereum event
	Unsupported Event = iota
	// MsgBurn is a Cosmos msg of type MsgBurn
	MsgBurn
	// MsgLock is a Cosmos msg of type MsgLock
	MsgLock
	// LogLock is an Ethereum event named 'LockEvent'
	LogLock
	// LogNewProphecyClaim is an Ethereum event named 'LogNewProphecyClaim'
	LogNewProphecyClaim
)

// String returns the event type as a string
func (d Event) String() string {
	return [...]string{"unsupported", "burn", "lock", "LogLock", "LogNewProphecyClaim"}[d]
}

// LockEvent struct which represents a LogLock event
type LockEvent struct {
	EthereumChainID       *big.Int
	BridgeContractAddress common.Address
	ID                    [32]byte
	From                  common.Address
	To                    []byte
	Token                 common.Address
	Symbol                string
	Value                 *big.Int
	Nonce                 *big.Int
}

// NewLockEvent creates a new LockEvent
func NewLockEvent(ethereumChainID *big.Int, bridgeContractAddress common.Address,
	id [32]byte, from common.Address, to []byte, token common.Address, symbol string,
	value *big.Int, nonce *big.Int) LockEvent {
	return LockEvent{
		EthereumChainID:       ethereumChainID,
		BridgeContractAddress: bridgeContractAddress,
		ID:                    id,
		From:                  from,
		To:                    to,
		Token:                 token,
		Symbol:                symbol,
		Value:                 value,
		Nonce:                 nonce,
	}
}

// String implements fmt.Stringer
func (l LockEvent) String() string {
	return fmt.Sprintf("\nChain ID: %v\nBridge contract address: %v\nToken symbol: %v\nToken"+
		"contract address: %v\nSender: %v\nRecipient: %v\nValue: %v\nNonce: %v\n\n",
		l.EthereumChainID, l.BridgeContractAddress.Hex(), l.Symbol, l.Token.Hex(), l.From.Hex(),
		string(l.To), l.Value, l.Nonce)
}

// ProphecyClaimEvent struct which represents a LogNewProphecyClaim event
type ProphecyClaimEvent struct {
	CosmosSender     []byte
	Symbol           string
	ProphecyID       *big.Int
	Amount           *big.Int
	EthereumReceiver common.Address
	ValidatorAddress common.Address
	TokenAddress     common.Address
	ClaimType        uint8
}

// NewProphecyClaimEvent creates a new ProphecyClaimEvent
func NewProphecyClaimEvent(cosmosSender []byte, symbol string, prophecyID, amount *big.Int, ethereumReceiver,
	validatorAddress, tokenAddress common.Address, claimType uint8) ProphecyClaimEvent {
	return ProphecyClaimEvent{
		CosmosSender:     cosmosSender,
		Symbol:           symbol,
		ProphecyID:       prophecyID,
		Amount:           amount,
		EthereumReceiver: ethereumReceiver,
		ValidatorAddress: validatorAddress,
		TokenAddress:     tokenAddress,
		ClaimType:        claimType,
	}
}

// String implements fmt.Stringer
func (p ProphecyClaimEvent) String() string {
	return fmt.Sprintf("\nProphecy ID: %v\nClaim Type: %v\nSender: %v\n"+
		"Recipient: %v\nSymbol %v\nToken %v\nAmount: %v\nValidator: %v\n\n",
		p.ProphecyID, p.ClaimType, string(p.CosmosSender), p.EthereumReceiver.Hex(),
		p.Symbol, p.TokenAddress.Hex(), p.Amount, p.ValidatorAddress.Hex())
}

// CosmosMsg contains data from MsgBurn and MsgLock events
type CosmosMsg struct {
	ClaimType            Event
	CosmosSender         []byte
	EthereumReceiver     common.Address
	TokenContractAddress common.Address
	Symbol               string
	Amount               *big.Int
}

// NewCosmosMsg creates a new CosmosMsg
func NewCosmosMsg(claimType Event, cosmosSender []byte, ethereumReceiver common.Address, symbol string,
	amount *big.Int, tokenContractAddress common.Address) CosmosMsg {
	return CosmosMsg{
		ClaimType:            claimType,
		CosmosSender:         cosmosSender,
		EthereumReceiver:     ethereumReceiver,
		Symbol:               symbol,
		Amount:               amount,
		TokenContractAddress: tokenContractAddress,
	}
}

// String implements fmt.Stringer
func (c CosmosMsg) String() string {
	return fmt.Sprintf("\nClaim Type: %v\nCosmos Sender: %v\nEthereum Recipient: %v"+
		"\nToken Address: %v\nSymbol: %v\nAmount: %v\n",
		c.ClaimType.String(), string(c.CosmosSender), c.EthereumReceiver.Hex(),
		c.TokenContractAddress.Hex(), c.Symbol, c.Amount)
}

// CosmosMsgAttributeKey enum containing supported attribute keys
type CosmosMsgAttributeKey int

const (
	// UnsupportedAttributeKey unsupported attribute key
	UnsupportedAttributeKey CosmosMsgAttributeKey = iota
	// CosmosSender sender's address on Cosmos network
	CosmosSender
	// EthereumReceiver receiver's address on Ethereum network
	EthereumReceiver
	// Coin is the coin type
	Coin
	// TokenContractAddress coin's corresponding contract address deployed on the Ethereum network
	TokenContractAddress
)

// String returns the event type as a string
func (d CosmosMsgAttributeKey) String() string {
	return [...]string{"unsupported", "cosmos_sender", "ethereum_receiver", "amount", "token_contract_address"}[d]
}
