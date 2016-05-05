package reader

import (
    "github.com/streadway/amqp"
    "github.com/ryanbrushett/msg-worker/common"
    "os"
)

func Consume() (*amqp.Connection, <-chan amqp.Delivery, error) {
    d := os.Getenv("MSG_CONFIG_HOME")
    if len(d) < 1 {
        panic("Please Set MSG_CONFIG_HOME variable")
    }
    if d[len(d)-1:] != "/" {
        d = common.Concat(d,"/")
    }
    f := "config.json"
    p := common.PropertiesJson(d,f)
    h := common.AMQPConnectionString(p)
    q := p.AckQueue
    v := p.VirtHost
    conn, dialErr := amqp.Dial(h)
    common.CheckError(dialErr)

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

    messages, err := c.Consume(q, q, false, false, false, false, nil)
    return conn, messages, err
}