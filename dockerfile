# Official Ollama image
FROM ollama/ollama

# Skip the time zone prompt
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies for downloading Go
RUN apt-get update && \
    apt-get install -y curl tar && \
    apt-get clean

# Install Go
ENV GOLANG_VERSION=1.24.0
RUN curl -LO https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz

# Set Go paths
ENV PATH="/usr/local/go/bin:${PATH}"

# Set working directory
WORKDIR /app

# Copy
COPY . .

# Launch ollama daemon
RUN ollama serve

# Pull a model
RUN ollama run gemma:2b

# Build binary
RUN go build -o load-balanced-llm .

# Expose Ollama’s default port
EXPOSE 11434

# Override ollama's default ENTRYPOINT
ENTRYPOINT []

# Run ollama and execute binary
CMD ["sh", "-c", "ollama serve & ./load-balanced-llm"]
