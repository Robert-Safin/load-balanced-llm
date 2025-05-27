package context

import (
	"Robert-Safin/load-balanced-llm/utils"
	"os"
)

func Load_context() (string, string) {
	dat, err := os.ReadFile("./context/context.txt")
	utils.Check(err, "could not find file")
	context := string(dat)

	dat, err = os.ReadFile("./context/prelude.txt")
	utils.Check(err, "could not find file")
	prelude := string(dat)

	return context, prelude
}
