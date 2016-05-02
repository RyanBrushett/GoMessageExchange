package main

import (
    "github.com/ryanbrushett/msg-worker/common"
    "github.com/ryanbrushett/msg-worker/writer"
    "strconv"
)

func main() {
    for i := 0; i < 100; i++ {
        message := common.Concat("This is a message: ", strconv.Itoa(i))
        err := writer.Write(message, "test-queue", "processing@rabbitmq.net")
        common.CheckError(err)
    }
}
