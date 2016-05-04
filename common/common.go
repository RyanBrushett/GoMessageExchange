package common

import (
    "bytes"
    "encoding/json"
    "os"
)

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func Concat(strings ...string) string {
    var b bytes.Buffer
    for _,s := range strings {
        b.WriteString(s)
    }
    return b.String()
}

func PropertiesJson(d,f string) Properties {
    p,e := os.Open(Concat(d,f))
    CheckError(e)
    jp := json.NewDecoder(p)
    properties := Properties{}
    CheckError(jp.Decode(&properties))
    return properties
}

func AMQPConnectionString(p Properties) string {
    return Concat(
        "amqp://",
        p.Username,":",p.Password,"@",
        p.Hostname,":",p.AMQPport,"/",
        p.VirtHost,
    )
}