package grifts

import (
	"github.com/adefemi171/blog/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
