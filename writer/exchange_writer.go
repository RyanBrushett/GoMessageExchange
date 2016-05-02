package writer

import (
    "github.com/streadway/amqp"
    "github.com/ryanbrushett/msg-worker/common"
    "time"
)

func Write(m string, q string, v string) error {
    host := common.Concat("amqp://guest:guest@127.0.0.1:5672/",v)
    conn, dialErr := amqp.Dial(host)
    common.CheckError(dialErr)
    defer conn.Close()

    c, channelErr := conn.Channel()
    common.CheckError(channelErr)

    declareErr := c.ExchangeDeclare(v,"topic",true,false,false,false,nil)
    common.CheckError(declareErr)

    message := amqp.Publishing{
        DeliveryMode: amqp.Persistent,
        Timestamp:    time.Now(),
        ContentType:  "text/plain",
        Body:         []byte(m),
    }

    publishErr := c.Publish(v,q,false,false,message)
    return publishErr
}