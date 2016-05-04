package main

import (
    "github.com/ryanbrushett/msg-worker/reader"
    "github.com/ryanbrushett/msg-worker/common"
)

func main() {
    p := common.PropertiesJson("/Users/ryan/Documents/code/ryanbrushett/msg-worker/properties/","config.json")
    rmq := common.AMQPConnectionString(p)
    reader.Read(rmq, p.AckQueue, p.VirtHost)
}
