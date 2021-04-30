package service

import (
	"ahobson/gofutz/thing"
)

type Service struct {
	thing thing.Thing
}

func NewService(thing thing.Thing) (Service) {
	return Service{
		thing: thing,
	}
}

func (s *Service) DoTheThing(id int) string {
	return s.thing.FetchThing(thing.ThingInput{Id: id}).Name
}
