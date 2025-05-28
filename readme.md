# About

This is a Go implementation for running Ollama LLMs locally. The generic use is a support chatbot. The model can be provided with system instructions, which alters its general behavior. The model will be preloaded with user defined context, which it will use to answer questions about said context. The project does not scaffold a HTTP server or any auth.

# Set up (No Docker)
Install [go toolchain](https://go.dev/) and [ollama](https://ollama.com/) with your choice of ollama LLM model.

Sample model install:

    ollama pull gemma:2b

To find out your model's name, run:

    ollama list

### Running go directly (test running)
    ollama serve
    go run . -prompt="prompt" -model="model name"

### Running binary (with context modifications)
Execute provided binary directly:

    ./load-balanced-llm -prompt="prompt" -model="model name"

Or compile youself with fresh context and system instructions:

    go build .
    ./load-balanced-llm -prompt="prompt" -model="model name"


# Set up (Docker)
The dockerfile uses official ollama image (~2gb) and install any dependencies required to compile a fresh binary and execute it. Though of course its not ideal to build a new container each time.

# The moving parts

### CLI Arguements
The command line will expect 2 mandatory arguments: prompt & model.

### Model params

`ChatRequest` struct contains some other unused params for modifying model behaviour.

### Context and Prelude
`prelude.txt` is used to give model system level instructions on how to behave.
`context.txt` is the user provided knowledge, the model is expected to know and respond with.



# Motivation
Python (unfortunately) is dominant language for ML/AI. Ollama thankfully provides as a nice API to interract with LLMs without touching python (fortunately). Although ollama can receive concurrent requests via queuing it is not able to process concurrently (thanks python).

Spinning up multiple instances of the container with a load balancer in front will solve this concurrency issue, leading to greater scalability.


# Going Further
With introduction of a HTTP server and auth, this can be quickly integrated as a sorts of customer support bot that is knowledgeable in ones particular use case and integrated into any form of HTTP based communication: whatsapp, telegram, discord, fb messenger, instagram, slack, email.
