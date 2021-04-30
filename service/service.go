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

func (s *Service) DoTheThing(id int) (string, error) {
	output, err := s.thing.FetchThing(thing.ThingInput{Id: id})
	if err != nil {
		return "", err
	}
	return output.Name, nil
}
