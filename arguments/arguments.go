package arguments

import (
	"Robert-Safin/load-balanced-llm/utils"
	"errors"
	"flag"
)

func ReadArguments() (string, string) {
	prompt := flag.String("prompt", "", "Prompt")

	model_name := flag.String("model", "llama3.2:latest", "Model name")

	flag.Parse()

	if *prompt == "" {
		utils.Check(errors.New("err"), "prompt is required")
	}
	return *prompt, *model_name
}
