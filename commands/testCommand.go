package commands

import (
	"github.com/Lukaesebrot/dgc"
)

func testCommand(ctx *dgc.Ctx) {
	ctx.RespondText("working !")
}
