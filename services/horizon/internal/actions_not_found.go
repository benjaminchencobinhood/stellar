package horizon

import (
	"github.com/benjaminchencobinhood/stellar/support/render/problem"
)

// NotFoundAction renders a 404 response
type NotFoundAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotFoundAction) JSON() {
	problem.Render(action.R.Context(), action.W, problem.NotFound)
}
