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
	FetchThing(thingInput ThingInput) ThingOutput
}

type RealThing struct {
}

func (rt *RealThing) FetchThing(thingInput ThingInput) ThingOutput {
	if thingInput.Id < 0 {
		return ThingOutput{
			Name: fmt.Sprintf("WILD THING: %d", thingInput.Id),
		}
	}
	return ThingOutput{
		Name: fmt.Sprintf("Thing Name: %d", thingInput.Id),
	}
}
