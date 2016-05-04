package main

import (
    "github.com/ryanbrushett/msg-worker/common"
    "github.com/ryanbrushett/msg-worker/writer"
    "strconv"
)

func main() {
    p := common.PropertiesJson("/Users/ryan/Documents/code/ryanbrushett/msg-worker/properties/","config.json")
    rmq := common.AMQPConnectionString(p)
    for i := 0; i < 100; i++ {
        message := common.Concat("This is a message: ", strconv.Itoa(i))
        common.CheckError(
            writer.Write(message, rmq, p.VirtHost, p.AckQueue),
        )
    }
}
