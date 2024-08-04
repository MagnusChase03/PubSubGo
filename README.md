# PubSubGo

**This is a simple pub/sub service written in Go.**

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

### NewPublisher(id int) *Publisher

Returns a new publisher with given ID.

### Subscribe(s *Subscriber, topic string)

Subscribes the subscriber to recieve messages from given topic.

### Unsubscribe(s *Subscriber, topic string)

Unsubscribes the subscriber from a given topic.

### Publish(topic string, message string)

Sends out a given message to all subscribers in a given topic.

### Subscriber

A subscriber is responsible for handling incoming messages from subscribed topics.

```go
type Subscriber struct {
    Id int
    Messages chan *Message
    Topics map[string]bool
    Active bool
    Mutex sync.RWMutex
}
```

### NewSubscriber(id int) *Subscriber

Returns a new subscriber with given ID.

### AddTopic(topic string)

Marks a topic as being subscriber to.

### RemoveTopic(topic string)

Marks a topic as not subscribed to.

### Signal(message *Message)

Adds a new message to the queue of messages to be processed.

### Listen()

Continously processes new messages in the queue.
