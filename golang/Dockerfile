FROM golang:1.25.5-alpine

# Install dependencies
RUN apk add --no-cache git ca-certificates gcc musl-dev

WORKDIR /app

# Copy source code
COPY src/ ./

# Initialize or download modules
RUN if [ ! -f go.mod ]; then \
      go mod init myapp; \
    else \
      go mod download; \
    fi

CMD ["tail", "-f", "/dev/null"]