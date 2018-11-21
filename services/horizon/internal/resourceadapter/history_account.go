package resourceadapter

import (
	"context"

	"github.com/benjaminchencobinhood/stellar/services/horizon/internal/db2/history"
	. "github.com/benjaminchencobinhood/stellar/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
