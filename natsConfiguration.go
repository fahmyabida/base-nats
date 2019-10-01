package main

import (
	"time"

	"github.com/nats-io/nats.go"
)

type natsConfiguration Service

// NatsMachine is instace variable to represent Nats object behavior
var (
	NatsEngine INats
	NatsConfig natsConfiguration
)

// Nats is used to create object nats
type Nats struct {
	natsEncode *nats.EncodedConn
}

// Publish is used to create command to publish
func (n Nats) Publish(topic string, data interface{}) {
	n.natsEncode.Publish(topic, data)
}

// Subscribe is used to create command to subcribe
func (n Nats) Subscribe(topic string, field interface{}) {
	n.natsEncode.Subscribe(topic, field)
}

// Request is used to create command to subcribe
func (n Nats) Request(topic string, field interface{}, reply interface{}, timeout time.Duration) {
	n.natsEncode.Request(topic, field, reply, timeout)
}

// SubscribeRequest is used to create command to reply a request
func (n Nats) SubscribeRequest(topic string, field interface{}, reply string) {
	n.natsEncode.PublishRequest(topic, reply, field)
}
