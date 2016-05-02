package reader

import (
    "github.com/streadway/amqp"
    "github.com/ryanbrushett/msg-worker/common"
    "runtime"
    "fmt"
    "time"
)

func Read(q string, v string) {
    conn, dialErr := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/" + v)
    common.CheckError(dialErr)
    defer conn.Close()

    c, channelError := conn.Channel()
    common.CheckError(channelError)

    exchangeErr := c.ExchangeDeclare(v,"topic",true,false,false,false,nil)
    common.CheckError(exchangeErr)

    _,queueErr := c.QueueDeclare(q,true,false,false,false,nil)
    common.CheckError(queueErr)

    bindErr := c.QueueBind(q,q,v,false,nil)
    common.CheckError(bindErr)

    qosErr := c.Qos(8,0,false)
    common.CheckError(qosErr)

    messages,consumeErr := c.Consume(q,q,false,false,false,false,nil)
    common.CheckError(consumeErr)

    for i := 0; i < runtime.NumCPU(); i++ {
        go func(work <-chan amqp.Delivery) {
            for message := range work {
                fmt.Println(string(message.Body))
                message.Ack(false)
            }
        }(messages)
    }

    time.Sleep(10 * time.Second)
}