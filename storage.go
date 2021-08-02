package main

type Storage interface {
	Add(Event)
}
