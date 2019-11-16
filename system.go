package starsystem

import (
	"context"
	"path/filepath"
)

type System struct {
	Name   string
	Actors map[string]*ActorProp
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSystem(name string) *System {
	ctx, cancel := context.WithCancel(context.Background())
	return &System{
		Name:   name,
		Actors: map[string]*ActorProp{},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *System) ActorOf(name string, val Actor) *ActorProp {
	_, ok := interface{}(val).(Actor) // test does val implement Actor interfce
	if ok {
		key := join(s.Name, name)
		prop := NewActorProp(key, s)
		prop.start(s.ctx, val)
		s.Actors[key] = prop
		return prop
	}

	return nil
}

func (s *System) watch(parrent *ActorProp) {
	go func() {
		for {
			select {
			case path := <-parrent.watch:
				name := filepath.Base(path)
				parrent.Tell(name)
			case <-s.ctx.Done():
				return
			}
		}
	}()
}

func (s *System) Stop(path string) bool {
	if val := s.ActorSelection(path); val != nil {
		val.notify()
		return true
	}
	return false
}

func (s *System) ActorSelection(path string) *ActorProp {
	if val, ok := s.Actors[path]; ok {
		return val
	}
	return nil
}

func (s *System) Shutdown() {
	s.cancel()
}
