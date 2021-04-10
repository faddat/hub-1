// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/node/types
// ALIASGEN: github.com/sentinel-official/hub/x/node/keeper
package node

import (
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

const (
	AttributeKeyProvider    = types.AttributeKeyProvider
	AttributeKeyAddress     = types.AttributeKeyAddress
	AttributeKeyStatus      = types.AttributeKeyStatus
	ModuleName              = types.ModuleName
	QuerierRoute            = types.QuerierRoute
	DefaultInactiveDuration = types.DefaultInactiveDuration
)

var (
	// functions aliases
	RegisterLegacyAminoCodec                = types.RegisterLegacyAminoCodec
	RegisterInterfaces                      = types.RegisterInterfaces
	NewGenesisState                         = types.NewGenesisState
	DefaultGenesisState                     = types.DefaultGenesisState
	NodeKey                                 = types.NodeKey
	ActiveNodeKey                           = types.ActiveNodeKey
	InactiveNodeKey                         = types.InactiveNodeKey
	GetActiveNodeForProviderKeyPrefix       = types.GetActiveNodeForProviderKeyPrefix
	ActiveNodeForProviderKey                = types.ActiveNodeForProviderKey
	GetInactiveNodeForProviderKeyPrefix     = types.GetInactiveNodeForProviderKeyPrefix
	InactiveNodeForProviderKey              = types.InactiveNodeForProviderKey
	GetActiveNodeAtKeyPrefix                = types.GetActiveNodeAtKeyPrefix
	ActiveNodeAtKey                         = types.ActiveNodeAtKey
	GetInactiveNodeAtKeyPrefix              = types.GetInactiveNodeAtKeyPrefix
	InactiveNodeAtKey                       = types.InactiveNodeAtKey
	AddressFromStatusNodeKey                = types.AddressFromStatusNodeKey
	AddressFromStatusNodeForProviderKey     = types.AddressFromStatusNodeForProviderKey
	AddressFromStatusNodeAtKey              = types.AddressFromStatusNodeAtKey
	NewMsgRegisterRequest                   = types.NewMsgRegisterRequest
	NewMsgUpdateRequest                     = types.NewMsgUpdateRequest
	NewMsgSetStatusRequest                  = types.NewMsgSetStatusRequest
	NewMsgServiceClient                     = types.NewMsgServiceClient
	RegisterMsgServiceServer                = types.RegisterMsgServiceServer
	NewParams                               = types.NewParams
	DefaultParams                           = types.DefaultParams
	ParamsKeyTable                          = types.ParamsKeyTable
	NewQueryNodeRequest                     = types.NewQueryNodeRequest
	NewQueryNodesRequest                    = types.NewQueryNodesRequest
	NewQueryNodesForProviderRequest         = types.NewQueryNodesForProviderRequest
	NewQueryServiceClient                   = types.NewQueryServiceClient
	RegisterQueryServiceServer              = types.RegisterQueryServiceServer
	RegisterQueryServiceHandlerServer       = types.RegisterQueryServiceHandlerServer
	RegisterQueryServiceHandlerFromEndpoint = types.RegisterQueryServiceHandlerFromEndpoint
	RegisterQueryServiceHandler             = types.RegisterQueryServiceHandler
	RegisterQueryServiceHandlerClient       = types.RegisterQueryServiceHandlerClient
	NewKeeper                               = keeper.NewKeeper
	NewMsgServiceServer                     = keeper.NewMsgServiceServer

	// variable aliases
	ModuleCdc                        = types.ModuleCdc
	ErrorMarshal                     = types.ErrorMarshal
	ErrorUnmarshal                   = types.ErrorUnmarshal
	ErrorUnknownMsgType              = types.ErrorUnknownMsgType
	ErrorUnknownQueryType            = types.ErrorUnknownQueryType
	ErrorInvalidField                = types.ErrorInvalidField
	ErrorProviderDoesNotExist        = types.ErrorProviderDoesNotExist
	ErrorDuplicateNode               = types.ErrorDuplicateNode
	ErrorNodeDoesNotExist            = types.ErrorNodeDoesNotExist
	EventTypeSet                     = types.EventTypeSet
	EventTypeUpdate                  = types.EventTypeUpdate
	EventTypeSetStatus               = types.EventTypeSetStatus
	ErrInvalidLengthGenesis          = types.ErrInvalidLengthGenesis
	ErrIntOverflowGenesis            = types.ErrIntOverflowGenesis
	ErrUnexpectedEndOfGroupGenesis   = types.ErrUnexpectedEndOfGroupGenesis
	ParamsSubspace                   = types.ParamsSubspace
	RouterKey                        = types.RouterKey
	StoreKey                         = types.StoreKey
	EventModuleName                  = types.EventModuleName
	NodeKeyPrefix                    = types.NodeKeyPrefix
	ActiveNodeKeyPrefix              = types.ActiveNodeKeyPrefix
	InactiveNodeKeyPrefix            = types.InactiveNodeKeyPrefix
	ActiveNodeForProviderKeyPrefix   = types.ActiveNodeForProviderKeyPrefix
	InactiveNodeForProviderKeyPrefix = types.InactiveNodeForProviderKeyPrefix
	ActiveNodeAtKeyPrefix            = types.ActiveNodeAtKeyPrefix
	InactiveNodeAtKeyPrefix          = types.InactiveNodeAtKeyPrefix
	ErrInvalidLengthMsg              = types.ErrInvalidLengthMsg
	ErrIntOverflowMsg                = types.ErrIntOverflowMsg
	ErrUnexpectedEndOfGroupMsg       = types.ErrUnexpectedEndOfGroupMsg
	ErrInvalidLengthNode             = types.ErrInvalidLengthNode
	ErrIntOverflowNode               = types.ErrIntOverflowNode
	ErrUnexpectedEndOfGroupNode      = types.ErrUnexpectedEndOfGroupNode
	DefaultDeposit                   = types.DefaultDeposit
	KeyDeposit                       = types.KeyDeposit
	KeyInactiveDuration              = types.KeyInactiveDuration
	ErrInvalidLengthParams           = types.ErrInvalidLengthParams
	ErrIntOverflowParams             = types.ErrIntOverflowParams
	ErrUnexpectedEndOfGroupParams    = types.ErrUnexpectedEndOfGroupParams
	ErrInvalidLengthQuerier          = types.ErrInvalidLengthQuerier
	ErrIntOverflowQuerier            = types.ErrIntOverflowQuerier
	ErrUnexpectedEndOfGroupQuerier   = types.ErrUnexpectedEndOfGroupQuerier
)

type (
	GenesisState                    = types.GenesisState
	MsgRegisterRequest              = types.MsgRegisterRequest
	MsgUpdateRequest                = types.MsgUpdateRequest
	MsgSetStatusRequest             = types.MsgSetStatusRequest
	MsgRegisterResponse             = types.MsgRegisterResponse
	MsgUpdateResponse               = types.MsgUpdateResponse
	MsgSetStatusResponse            = types.MsgSetStatusResponse
	MsgServiceClient                = types.MsgServiceClient
	MsgServiceServer                = types.MsgServiceServer
	UnimplementedMsgServiceServer   = types.UnimplementedMsgServiceServer
	Nodes                           = types.Nodes
	Node                            = types.Node
	Params                          = types.Params
	QueryNodesRequest               = types.QueryNodesRequest
	QueryNodesForProviderRequest    = types.QueryNodesForProviderRequest
	QueryNodeRequest                = types.QueryNodeRequest
	QueryNodesResponse              = types.QueryNodesResponse
	QueryNodesForProviderResponse   = types.QueryNodesForProviderResponse
	QueryNodeResponse               = types.QueryNodeResponse
	QueryServiceClient              = types.QueryServiceClient
	QueryServiceServer              = types.QueryServiceServer
	UnimplementedQueryServiceServer = types.UnimplementedQueryServiceServer
	Keeper                          = keeper.Keeper
	Querier                         = keeper.Querier
)
