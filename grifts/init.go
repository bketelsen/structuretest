package grifts

import (
	"github.com/bketelsen/structuretest/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
