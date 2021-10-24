package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "custommint"
	StoreKey     = ModuleName
	QuerierRoute = ModuleName
)

var (
	InflationKeyPrefix = []byte{0x01}
)

func InflationKey(timestamp time.Time) []byte {
	return append(InflationKeyPrefix, sdk.FormatTimeBytes(timestamp)...)
}
