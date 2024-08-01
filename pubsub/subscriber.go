package pubsub

import (
    "sync"
    "fmt"
)

type Subscriber struct {
    Id int
    Messages chan *Message
    Topics map[string]bool
    Active bool
    Mutex sync.RWMutex
}

func NewSubscriber(id int) *Subscriber {
    return &Subscriber{
        Id: id,
        Messages: make(chan *Message),
        Topics: map[string]bool{},
        Active: true,
    }
}

func (s *Subscriber) AddTopic(topic string) {
    s.Mutex.Lock()
    defer s.Mutex.Unlock()
    s.Topics[topic] = true

    fmt.Printf("Subscriber %d is subscribed to %s.\n", s.Id, topic)
}

func (s *Subscriber) Signal(message *Message) {
    s.Mutex.Lock()
    defer s.Mutex.Unlock()
    s.Messages <- message
}

func (s *Subscriber) Listen() {
    for {
        if m, ok := <- s.Messages; ok {
            fmt.Printf("Subscriber %d recieved [%s] %s.\n", s.Id, m.Topic, m.Body)
        }
    }
}
