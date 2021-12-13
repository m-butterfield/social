# Bases for building and running the app
FROM golang:1.17.4-alpine AS builder-base
WORKDIR /go/src/github.com/m-butterfield/social
COPY go.* ./
RUN go mod download
ADD . /go/src/github.com/m-butterfield/social

FROM alpine:latest AS runner-base
WORKDIR /root

# Run build
FROM builder-base AS server-builder
RUN go build -o bin/server cmd/server/main.go

# Copy the built executable to the runner
FROM runner-base AS server
COPY --from=server-builder /go/src/github.com/m-butterfield/social/bin/ ./bin/
CMD ["bin/server"]