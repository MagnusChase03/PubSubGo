package pubsub

import (
    "sync"
    "fmt"
)

type Publisher struct {
    Id int
    Subscribers map[int]*Subscriber
    Topics map[string]map[int]*Subscriber 
    Mutex sync.RWMutex
}

func NewPublisher(id int) *Publisher {
    return &Publisher{
        Id: id,
        Subscribers: map[int]*Subscriber{},
        Topics: map[string]map[int]*Subscriber{},
    } 
}

func (p *Publisher) Subscribe(s *Subscriber, topic string) {
    p.Mutex.Lock()
    defer p.Mutex.Unlock()

    if p.Topics[topic] == nil {
        p.Topics[topic] = map[int]*Subscriber{}
    }

    if p.Subscribers[s.Id] == nil {
        p.Subscribers[s.Id] = s
    }

    s.AddTopic(topic)
    p.Topics[topic][s.Id] = s
    fmt.Printf("Subscriber %d subscribed to %s via Publisher %d.\n", s.Id, topic, p.Id)
}

func (p *Publisher) Unsubscribe(s *Subscriber, topic string) {
    p.Mutex.Lock()
    defer p.Mutex.Unlock()

    s.RemoveTopic(topic)
    delete(p.Topics[topic], s.Id)
    fmt.Printf("Subscriber %d unsubscribed from %s via Publisher %d.\n", s.Id, topic, p.Id)
}

func (p *Publisher) Publish(topic string, message string) {
    p.Mutex.RLock()
    defer p.Mutex.RUnlock()
    topicSubs := p.Topics[topic]

    fmt.Printf("Publisher %d publishing [%s] %s.\n", p.Id, topic, message)
    m := NewMessage(topic, message)
    for _, s := range topicSubs {
        if !s.Active {
            continue
        } 

        go (func (s *Subscriber) {
            s.Signal(m)
        })(s)
    }
}
