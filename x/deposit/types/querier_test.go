package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestNewQueryDepositRequest(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		require.Equal(
			t,
			&QueryDepositRequest{
				Address: sdk.AccAddress(address).String(),
			},
			NewQueryDepositRequest(address),
		)
	}
}

func TestNewQueryDepositsRequest(t *testing.T) {
	var (
		pagination *query.PageRequest
	)

	for i := 0; i < 20; i++ {
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryDepositsRequest{
				Pagination: pagination,
			},
			NewQueryDepositsRequest(pagination),
		)
	}
}
