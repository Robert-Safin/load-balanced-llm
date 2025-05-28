package arguments

import (
	"Robert-Safin/load-balanced-llm/utils"
	"errors"
	"flag"
)

func ReadArguments() (string, string) {
	prompt := flag.String("prompt", "", "Prompt")

	model_name := flag.String("model", "", "Model name")

	flag.Parse()

	if *prompt == "" {
		utils.Check(errors.New("err"), "prompt is required")
	}

	if *model_name == "" {
		utils.Check(errors.New("err"), "model name is required")
	}
	return *prompt, *model_name
}
