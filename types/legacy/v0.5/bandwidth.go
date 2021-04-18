package v0_5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Bandwidth struct {
		Upload   sdk.Int `json:"upload"`
		Download sdk.Int `json:"download"`
	}
)
