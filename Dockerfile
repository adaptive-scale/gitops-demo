FROM golang:alpine

`WORKDIR /workspace
COPY app.go /workspace

ENTRYPOINT ["sh", "-c", "go run /workspace/app.go"]
