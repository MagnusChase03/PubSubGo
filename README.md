# PubSubGo

`This is a simple pub/sub service written in Go.`

## Usage

To run the example utilization:

```
$ go run main.go
```

## API

To utilize as a library, here is the documentation:

**Structs**

- Publisher
    - NewPublisher
    - Subscribe
    - Unsubscribe
    - Publish
- Subsciber
    - NewSubscriber
    - AddTopic
    - RemoveTopic
    - Signal
    - Listen

### Publisher

A publisher is responsible for sending messages to those subscribed to topics.

```go
type Publisher struct {
    Id int
    Subscribers map[int]*Subscriber
    Topics map[string]map[int]*Subscriber 
    Mutex sync.RWMutex
}
```

`NewPublisher(id int) *Publisher`

Returns a new publisher with given ID.

`Subscribe(s *Subscriber, topic string)`

Subscribes the subscriber to recieve messages from given topic.

`Unsubscribe(s *Subscriber, topic string)`

Unsubscribes the subscriber from a given topic.

`Publish(topic string, message string)`

Sends out a given message to all subscribers in a given topic.

### Subscriber
