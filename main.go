package main

import (
	"Robert-Safin/load-balanced-llm/arguments"
	"Robert-Safin/load-balanced-llm/chat"
	"Robert-Safin/load-balanced-llm/context"
)

func main() {
	prompt_input, model_name_input := arguments.ReadArguments()

	context, prelude := context.Load_context()

	chat.Chat_api(chat.Generate_model_args{
		Prompt:  prompt_input,
		Model:   model_name_input,
		Context: context,
		Prelude: prelude,
	})

}
