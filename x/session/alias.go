// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/session/types
// ALIASGEN: github.com/sentinel-official/hub/x/session/keeper
package session

import (
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

const (
	AttributeKeyCount               = types.AttributeKeyCount
	AttributeKeyID                  = types.AttributeKeyID
	AttributeKeySubscription        = types.AttributeKeySubscription
	AttributeKeyAddress             = types.AttributeKeyAddress
	ModuleName                      = types.ModuleName
	QuerierRoute                    = types.QuerierRoute
	DefaultInactiveDuration         = types.DefaultInactiveDuration
	DefaultProofVerificationEnabled = types.DefaultProofVerificationEnabled
)

var (
	// functions aliases
	RegisterLegacyAminoCodec                = types.RegisterLegacyAminoCodec
	RegisterInterfaces                      = types.RegisterInterfaces
	NewGenesisState                         = types.NewGenesisState
	DefaultGenesisState                     = types.DefaultGenesisState
	GetChannelKeyPrefix                     = types.GetChannelKeyPrefix
	ChannelKey                              = types.ChannelKey
	SessionKey                              = types.SessionKey
	GetSessionForSubscriptionKeyPrefix      = types.GetSessionForSubscriptionKeyPrefix
	SessionForSubscriptionKey               = types.SessionForSubscriptionKey
	GetSessionForNodeKeyPrefix              = types.GetSessionForNodeKeyPrefix
	SessionForNodeKey                       = types.SessionForNodeKey
	GetSessionForAddressKeyPrefix           = types.GetSessionForAddressKeyPrefix
	SessionForAddressKey                    = types.SessionForAddressKey
	GetActiveSessionForAddressKeyPrefix     = types.GetActiveSessionForAddressKeyPrefix
	ActiveSessionForAddressKey              = types.ActiveSessionForAddressKey
	GetActiveSessionAtKeyPrefix             = types.GetActiveSessionAtKeyPrefix
	ActiveSessionAtKey                      = types.ActiveSessionAtKey
	IDFromSessionForSubscriptionKey         = types.IDFromSessionForSubscriptionKey
	IDFromSessionForNodeKey                 = types.IDFromSessionForNodeKey
	IDFromSessionForAddressKey              = types.IDFromSessionForAddressKey
	IDFromActiveSessionAtKey                = types.IDFromActiveSessionAtKey
	NewMsgUpsertRequest                     = types.NewMsgUpsertRequest
	NewMsgServiceClient                     = types.NewMsgServiceClient
	RegisterMsgServiceServer                = types.RegisterMsgServiceServer
	NewParams                               = types.NewParams
	DefaultParams                           = types.DefaultParams
	ParamsKeyTable                          = types.ParamsKeyTable
	NewProof                                = types.NewProof
	NewQuerySessionRequest                  = types.NewQuerySessionRequest
	NewQuerySessionsRequest                 = types.NewQuerySessionsRequest
	NewQuerySessionsForSubscriptionRequest  = types.NewQuerySessionsForSubscriptionRequest
	NewQuerySessionsForNodeRequest          = types.NewQuerySessionsForNodeRequest
	NewQuerySessionsForAddressRequest       = types.NewQuerySessionsForAddressRequest
	NewQueryActiveSessionRequest            = types.NewQueryActiveSessionRequest
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
	ErrorSubscriptionDoesNotExit     = types.ErrorSubscriptionDoesNotExit
	ErrorInvalidSubscriptionStatus   = types.ErrorInvalidSubscriptionStatus
	ErrorUnauthorized                = types.ErrorUnauthorized
	ErrorQuotaDoesNotExist           = types.ErrorQuotaDoesNotExist
	ErrorInvalidChannel              = types.ErrorInvalidChannel
	ErrorFailedToVerifyProof         = types.ErrorFailedToVerifyProof
	EventTypeSetCount                = types.EventTypeSetCount
	EventTypeSetActive               = types.EventTypeSetActive
	EventTypeUpdate                  = types.EventTypeUpdate
	ErrInvalidLengthGenesis          = types.ErrInvalidLengthGenesis
	ErrIntOverflowGenesis            = types.ErrIntOverflowGenesis
	ErrUnexpectedEndOfGroupGenesis   = types.ErrUnexpectedEndOfGroupGenesis
	ParamsSubspace                   = types.ParamsSubspace
	RouterKey                        = types.RouterKey
	StoreKey                         = types.StoreKey
	EventModuleName                  = types.EventModuleName
	CountKey                         = types.CountKey
	ChannelKeyPrefix                 = types.ChannelKeyPrefix
	SessionKeyPrefix                 = types.SessionKeyPrefix
	SessionForSubscriptionKeyPrefix  = types.SessionForSubscriptionKeyPrefix
	SessionForNodeKeyPrefix          = types.SessionForNodeKeyPrefix
	SessionForAddressKeyPrefix       = types.SessionForAddressKeyPrefix
	ActiveSessionAtKeyPrefix         = types.ActiveSessionAtKeyPrefix
	ActiveSessionForAddressKeyPrefix = types.ActiveSessionForAddressKeyPrefix
	ErrInvalidLengthMsg              = types.ErrInvalidLengthMsg
	ErrIntOverflowMsg                = types.ErrIntOverflowMsg
	ErrUnexpectedEndOfGroupMsg       = types.ErrUnexpectedEndOfGroupMsg
	KeyInactiveDuration              = types.KeyInactiveDuration
	KeyProofVerificationEnabled      = types.KeyProofVerificationEnabled
	ErrInvalidLengthParams           = types.ErrInvalidLengthParams
	ErrIntOverflowParams             = types.ErrIntOverflowParams
	ErrUnexpectedEndOfGroupParams    = types.ErrUnexpectedEndOfGroupParams
	ErrInvalidLengthProof            = types.ErrInvalidLengthProof
	ErrIntOverflowProof              = types.ErrIntOverflowProof
	ErrUnexpectedEndOfGroupProof     = types.ErrUnexpectedEndOfGroupProof
	ErrInvalidLengthQuerier          = types.ErrInvalidLengthQuerier
	ErrIntOverflowQuerier            = types.ErrIntOverflowQuerier
	ErrUnexpectedEndOfGroupQuerier   = types.ErrUnexpectedEndOfGroupQuerier
	ErrInvalidLengthSession          = types.ErrInvalidLengthSession
	ErrIntOverflowSession            = types.ErrIntOverflowSession
	ErrUnexpectedEndOfGroupSession   = types.ErrUnexpectedEndOfGroupSession
)

type (
	GenesisState                         = types.GenesisState
	MsgUpsertRequest                     = types.MsgUpsertRequest
	MsgUpsertResponse                    = types.MsgUpsertResponse
	MsgServiceClient                     = types.MsgServiceClient
	MsgServiceServer                     = types.MsgServiceServer
	UnimplementedMsgServiceServer        = types.UnimplementedMsgServiceServer
	Params                               = types.Params
	Proof                                = types.Proof
	QuerySessionsRequest                 = types.QuerySessionsRequest
	QuerySessionsForSubscriptionRequest  = types.QuerySessionsForSubscriptionRequest
	QuerySessionsForNodeRequest          = types.QuerySessionsForNodeRequest
	QuerySessionsForAddressRequest       = types.QuerySessionsForAddressRequest
	QuerySessionRequest                  = types.QuerySessionRequest
	QueryActiveSessionRequest            = types.QueryActiveSessionRequest
	QuerySessionsResponse                = types.QuerySessionsResponse
	QuerySessionsForSubscriptionResponse = types.QuerySessionsForSubscriptionResponse
	QuerySessionsForNodeResponse         = types.QuerySessionsForNodeResponse
	QuerySessionsForAddressResponse      = types.QuerySessionsForAddressResponse
	QuerySessionResponse                 = types.QuerySessionResponse
	QueryActiveSessionResponse           = types.QueryActiveSessionResponse
	QueryServiceClient                   = types.QueryServiceClient
	QueryServiceServer                   = types.QueryServiceServer
	UnimplementedQueryServiceServer      = types.UnimplementedQueryServiceServer
	Sessions                             = types.Sessions
	Session                              = types.Session
	Keeper                               = keeper.Keeper
	Querier                              = keeper.Querier
)
