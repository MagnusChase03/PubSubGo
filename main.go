package main

import (
    "time"
    "github.com/MagnusChase03/pubsubgo/pubsub"
)

func main() {
    p := pubsub.NewPublisher(0)
    s := pubsub.NewSubscriber(0)

    p.Subscribe(s, "anime")
    p.Publish("anime", "New season")

    go s.Listen()
    time.Sleep(1 * time.Second)
}
