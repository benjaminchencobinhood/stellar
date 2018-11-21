package resourceadapter

import (
	"github.com/benjaminchencobinhood/stellar/services/horizon/internal/db2/core"
	. "github.com/benjaminchencobinhood/stellar/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
