package commands

import (
	"github.com/Lukaesebrot/dgc"
)

func TestCommand(ctx *dgc.Ctx) {
	ctx.RespondText("working !")
}
