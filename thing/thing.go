package thing

import (
	"fmt"
)

type ThingInput struct {
	Id int
}

type ThingOutput struct {
	Name string
}

//go:generate mockery --name=Thing
type Thing interface {
	FetchThing(thingInput ThingInput) (ThingOutput, error)
}

type RealThing struct {
}

func (rt *RealThing) FetchThing(thingInput ThingInput) (ThingOutput, error) {
	if thingInput.Id < 0 {
		return ThingOutput{}, fmt.Errorf("Bad id")
	}
	output := ThingOutput{
		Name: fmt.Sprintf("Thing Name: %d", thingInput.Id),
	}
	return output, nil
}
