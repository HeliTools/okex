package private

import (
	"github.com/HeliTools/okex/events"
	"github.com/HeliTools/okex/models/account"
	"github.com/HeliTools/okex/models/trade"
)

type (
	Account struct {
		Arg      *events.Argument   `json:"arg"`
		Balances []*account.Balance `json:"data"`
	}
	Position struct {
		Arg       *events.Argument    `json:"arg"`
		Positions []*account.Position `json:"data"`
	}
	BalanceAndPosition struct {
		Arg                 *events.Argument              `json:"arg"`
		BalanceAndPositions []*account.BalanceAndPosition `json:"data"`
	}
	Order struct {
		Arg    *events.Argument `json:"arg"`
		Orders []*trade.Order   `json:"data"`
	}
)
