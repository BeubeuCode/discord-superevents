package commands

import (
	"github.com/lus/dgc"
)

func TestCommand(ctx *dgc.Ctx) {
	ctx.RespondText("working !")
}
