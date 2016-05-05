package main

import (
    "github.com/ryanbrushett/msg-worker/common"
    "github.com/ryanbrushett/msg-worker/writer"
    "strconv"
)

func main() {
    configDir := "/Users/ryan/Documents/code/ryanbrushett/msg-worker/properties/"
    configFile := "config.json"
    for i := 0; i < 100; i++ {
        message := common.Concat("This is a message: ", strconv.Itoa(i))
        common.CheckError(
            writer.Write(message,configDir,configFile),
        )
    }
}
