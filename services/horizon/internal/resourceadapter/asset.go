package resourceadapter

import (
	"context"

	"github.com/benjaminchencobinhood/stellar/xdr"
	. "github.com/benjaminchencobinhood/stellar/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
