package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/lenfree/pos/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
