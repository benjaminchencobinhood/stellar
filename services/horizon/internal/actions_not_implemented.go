package horizon

import (
	hProblem "github.com/benjaminchencobinhood/stellar/services/horizon/internal/render/problem"
	"github.com/benjaminchencobinhood/stellar/support/render/problem"
)

// NotImplementedAction renders a NotImplemented prblem
type NotImplementedAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotImplementedAction) JSON() {
	problem.Render(action.R.Context(), action.W, hProblem.NotImplemented)
}
