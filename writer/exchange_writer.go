package writer

import (
    "github.com/streadway/amqp"
    "github.com/ryanbrushett/msg-worker/common"
    "time"
)

func Write(m,d,f string) error {
    p := common.PropertiesJson(d,f)
    h := common.AMQPConnectionString(p)
    q := p.AckQueue
    v := p.VirtHost
    conn, dialErr := amqp.Dial(h)
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