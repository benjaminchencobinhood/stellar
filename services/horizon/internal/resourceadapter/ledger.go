package resourceadapter

import (
	"context"
	"fmt"

	"github.com/benjaminchencobinhood/stellar/amount"
	"github.com/benjaminchencobinhood/stellar/services/horizon/internal/db2/history"
	"github.com/benjaminchencobinhood/stellar/services/horizon/internal/httpx"
	"github.com/benjaminchencobinhood/stellar/xdr"
	. "github.com/benjaminchencobinhood/stellar/protocols/horizon"
	"github.com/benjaminchencobinhood/stellar/support/render/hal"
)

func PopulateLedger(ctx context.Context, dest *Ledger, row history.Ledger) {
	dest.ID = row.LedgerHash
	dest.PT = row.PagingToken()
	dest.Hash = row.LedgerHash
	dest.PrevHash = row.PreviousLedgerHash.String
	dest.Sequence = row.Sequence
	dest.TransactionCount = row.TransactionCount
	dest.OperationCount = row.OperationCount
	dest.ClosedAt = row.ClosedAt
	dest.TotalCoins = amount.String(xdr.Int64(row.TotalCoins))
	dest.FeePool = amount.String(xdr.Int64(row.FeePool))
	dest.BaseFee = row.BaseFee
	dest.BaseReserve = row.BaseReserve
	dest.MaxTxSetSize = row.MaxTxSetSize
	dest.ProtocolVersion = row.ProtocolVersion

	if row.LedgerHeaderXDR.Valid {
		dest.HeaderXDR = row.LedgerHeaderXDR.String
	} else {
		dest.HeaderXDR = ""
	}

	self := fmt.Sprintf("/ledgers/%d", row.Sequence)
	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Self = lb.Link(self)
	dest.Links.Transactions = lb.PagedLink(self, "transactions")
	dest.Links.Operations = lb.PagedLink(self, "operations")
	dest.Links.Payments = lb.PagedLink(self, "payments")
	dest.Links.Effects = lb.PagedLink(self, "effects")

	return
}