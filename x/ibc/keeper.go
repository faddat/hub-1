package ibc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func EgressKey(destChainID string, length int64) string {
	return fmt.Sprintf("egress/%s/%d", destChainID, length)
}

func EgressLengthKey(destChainID string) string {
	return fmt.Sprintf("egress/%s", destChainID)
}

func IngressKey(srcChainID string, length int64) string {
	return fmt.Sprintf("ingress/%s/%d", srcChainID, length)
}

func IngressLengthKey(srcChainID string) string {
	return fmt.Sprintf("ingress/%s", srcChainID)
}

type Keeper struct {
	IBCKey csdkTypes.StoreKey
	cdc    *codec.Codec
}

func NewKeeper(ibcKey csdkTypes.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		IBCKey: ibcKey,
		cdc:    cdc,
	}
}

func (k Keeper) SetIBCPacket(ctx csdkTypes.Context, packetID string, packet sdkTypes.IBCPacket) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryBare(packetID)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryBare(packet)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetIBCPacket(ctx csdkTypes.Context, packetID string) *sdkTypes.IBCPacket {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryBare(packetID)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil
	}

	var packet sdkTypes.IBCPacket

	if err := k.cdc.UnmarshalBinaryBare(valueBytes, &packet); err != nil {
		panic(err)
	}

	return &packet
}

func (k Keeper) SetEgressLength(ctx csdkTypes.Context, egressKey string, length int64) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryBare(egressKey)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryBare(length)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetEgressLength(ctx csdkTypes.Context, egressKey string) int64 {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryBare(egressKey)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0
	}

	var length int64

	if err := k.cdc.UnmarshalBinaryBare(valueBytes, &length); err != nil {
		panic(err)
	}

	return length
}

func (k Keeper) SetIngressLength(ctx csdkTypes.Context, ingressKey string, length int64) {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryBare(ingressKey)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryBare(length)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetIngressLength(ctx csdkTypes.Context, ingressKey string) int64 {
	store := ctx.KVStore(k.IBCKey)
	keyBytes, err := k.cdc.MarshalBinaryBare(ingressKey)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0
	}

	var length int64

	if err := k.cdc.UnmarshalBinaryBare(valueBytes, &length); err != nil {
		panic(err)
	}

	return length
}

func (k Keeper) PostIBCPacket(ctx csdkTypes.Context, packet sdkTypes.IBCPacket) csdkTypes.Error {
	egressLength := k.GetEgressLength(ctx, EgressLengthKey(packet.DestChainID))
	k.SetIBCPacket(ctx, EgressKey(packet.DestChainID, egressLength), packet)
	k.SetEgressLength(ctx, EgressLengthKey(packet.DestChainID), egressLength+1)

	return nil
}
