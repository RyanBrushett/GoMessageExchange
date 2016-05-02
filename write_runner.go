package main

import (
    "github.com/ryanbrushett/msg-worker/common"
    "github.com/ryanbrushett/msg-worker/writer"
)

func main() {
    err := writer.Write("This is a message", "test-queue", "processing@rabbitmq.net")
    common.CheckError(err)
}
