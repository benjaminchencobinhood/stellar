package horizon

import (
	"github.com/benjaminchencobinhood/stellar/protocols/horizon"
	"github.com/benjaminchencobinhood/stellar/services/horizon/internal/paths"
	"github.com/benjaminchencobinhood/stellar/services/horizon/internal/resourceadapter"
	"github.com/benjaminchencobinhood/stellar/support/render/hal"
)

// PathIndexAction provides path finding
type PathIndexAction struct {
	Action
	Query   paths.Query
	Records []paths.Path
	Page    hal.BasePage
}

// JSON implements actions.JSON
func (action *PathIndexAction) JSON() {
	action.Do(
		action.loadQuery,
		action.loadSourceAssets,
		action.loadRecords,
		action.loadPage,
		func() {
			hal.Render(action.W, action.Page)
		},
	)
}

func (action *PathIndexAction) loadQuery() {
	action.Query.DestinationAmount = action.GetPositiveAmount("destination_amount")
	action.Query.DestinationAddress = action.GetAddress("destination_account")
	action.Query.DestinationAsset = action.GetAsset("destination_")
}

func (action *PathIndexAction) loadSourceAssets() {
	app := AppFromContext(action.R.Context())
	protocolVersion := app.protocolVersion

	action.Err = action.CoreQ().AssetsForAddress(
		&action.Query.SourceAssets,
		action.GetAddress("source_account"),
		protocolVersion,
	)
}

func (action *PathIndexAction) loadRecords() {
	action.Records, action.Err = action.App.paths.Find(action.Query, action.App.config.MaxPathLength)
}

func (action *PathIndexAction) loadPage() {
	action.Page.Init()
	for _, p := range action.Records {
		var res horizon.Path
		action.Err = resourceadapter.PopulatePath(action.R.Context(), &res, action.Query, p)

		if action.Err != nil {
			return
		}
		action.Page.Add(res)
	}
}
