package common

type Properties struct{
    Hostname string `json:"hostname"`
    AMQPport string `json:"amqpport"`
    Username string `json:"username"`
    Password string `json:"password"`
    AckQueue string `json:"ackqueue"`
    VirtHost string `json:"virthost"`
}
