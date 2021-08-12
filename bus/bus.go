package bus

import (
	"log"

	"github.com/amirrezaask/connect/domain"
	"go.uber.org/zap"
)

type Bus interface {
	Emit(e *domain.Event) error
	// registering handlers should also be in the interface
	Register(t domain.EventType, handler func(e *domain.Event) error)
}

type ChannelBus struct {
	c      chan *domain.Event
	Logger *zap.SugaredLogger
	// maps event_types to their respective handlers
	handlers map[domain.EventType]func(e *domain.Event) error
}

func NewChannelBus() *ChannelBus {
	c := make(chan *domain.Event)
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	return &ChannelBus{
		c:        c,
		Logger:   l.Sugar(),
		handlers: map[domain.EventType]func(e *domain.Event) error{},
	}
}

func (c *ChannelBus) Emit(e *domain.Event) error {
	c.c <- e
	return nil
}

func (c *ChannelBus) Register(t domain.EventType, handler func(e *domain.Event) error) {
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
