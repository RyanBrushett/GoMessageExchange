package main

import (
    "github.com/ryanbrushett/msg-worker/reader"
    "github.com/ryanbrushett/msg-worker/common"
    "github.com/streadway/amqp"
    "fmt"
)

func doWork(work <-chan amqp.Delivery) {
    for message := range work {
        fmt.Println(string(message.Body))
        message.Ack(false)
    }
}

func main() {
    conn, messages, consumeErr := reader.Read("/Users/ryan/Documents/code/ryanbrushett/msg-worker/properties/","config.json")
    common.CheckError(consumeErr)
    forever := make(chan bool)
    go doWork(messages)
    <-forever
    conn.Close()
}
