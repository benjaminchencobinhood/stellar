package protocols

import (
	"fmt"

	"github.com/benjaminchencobinhood/stellar/build"
	shared "github.com/benjaminchencobinhood/stellar/services/internal/bridge-compliance-shared"
	"github.com/benjaminchencobinhood/stellar/support/errors"
)

// ToBaseAsset transforms Asset to github.com/benjaminchencobinhood/stellar-stellar-base/build.Asset
func (a Asset) ToBaseAsset() build.Asset {
	if a.Code == "" && a.Issuer == "" {
		return build.NativeAsset()
	}
	return build.CreditAsset(a.Code, a.Issuer)
}

// String returns string representation of this asset
func (a Asset) String() string {
	return fmt.Sprintf("Code: %s, Issuer: %s", a.Code, a.Issuer)
}

// Validate checks if asset params are correct.
func (a Asset) Validate() error {
	if a.Code != "" && a.Issuer != "" {
		if !shared.IsValidAssetCode(a.Code) {
			return errors.New("Invalid asset_code")
		}
		if !shared.IsValidAccountID(a.Issuer) {
			return errors.New("Invalid asset_issuer")
		}
	} else if a.Code == "" && a.Issuer == "" {
		// Native
		return nil
	} else {
		return errors.New("Asset code or issuer is missing")
	}

	return nil
}
