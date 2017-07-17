package grifts

import (
	"github.com/authdemo/githubauth/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
