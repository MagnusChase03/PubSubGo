package pubsub

type Message struct {
    Topic string
    Body string
}

func NewMessage(topic string, message string) *Message {
    return &Message{
        Topic: topic,
        Body: message,
    }
}
