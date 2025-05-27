package chat

import (
	"Robert-Safin/load-balanced-llm/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"net/http"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model         string        `json:"model"`
	Messages      []ChatMessage `json:"messages"`
	Stream        bool          `json:"stream"`
	Temperature   float64       `json:"temperature,omitempty"`
	TopP          float64       `json:"top_p,omitempty"`
	TopK          int           `json:"top_k,omitempty"`
	RepeatPenalty float64       `json:"repeat_penalty,omitempty"`
	Stop          []string      `json:"stop,omitempty"`
	NumPredict    int           `json:"num_predict,omitempty"`
}

type ChatResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	DoneReason         string `json:"done_reason"`
	Done               bool   `json:"done"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

type Generate_model_args struct {
	Prompt  string
	Model   string
	Context string
	Prelude string
}

func Chat_api(args Generate_model_args) {
	payload := ChatRequest{
		Model:  args.Model,
		Stream: false,
		Messages: []ChatMessage{
			{Role: "system", Content: args.Prelude + args.Context},
			{Role: "user", Content: args.Prompt},
		},
	}

	encoded, err := json.Marshal(payload)
	utils.Check(err, "failed to encode into json")

	res, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(encoded))

	utils.Check(err, "failed to hit model endpoint")
	defer res.Body.Close()

	var response ChatResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	utils.Check(err, "failed to decode response body")
	fmt.Println(response.Message.Content)

}
