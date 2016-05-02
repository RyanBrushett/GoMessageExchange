package main

import (
    "github.com/ryanbrushett/msg-worker/reader"
)

func main() {
    reader.Read("test-queue","processing@rabbitmq.net")
}
