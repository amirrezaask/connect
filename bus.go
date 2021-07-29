package main

import (
	"log"

	"go.uber.org/zap"
)

type Bus interface {
	Emit(e *Event) error
	// registering handlers should also be in the interface
	Register(t EventType, handler func(e *Event) error)
}

type ChannelBus struct {
	c      chan *Event
	Logger *zap.SugaredLogger
	// maps event_types to their respective handlers
	handlers map[EventType]func(e *Event) error
}

func NewChannelBus() *ChannelBus {
	c := make(chan *Event)
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	return &ChannelBus{
		c:        c,
		Logger:   l.Sugar(),
		handlers: map[EventType]func(e *Event) error{},
	}
}

func (c *ChannelBus) Emit(e *Event) error {
	c.c <- e
	return nil
}

func (c *ChannelBus) Register(t EventType, handler func(e *Event) error) {
	c.handlers[t] = handler
	go func() {
		for {
			e := <-c.c
			c.Logger.Debugf("handling :%+v", e)
			go c.handlers[e.EventType](e)
		}
	}()
}

type NATSBus struct{}
