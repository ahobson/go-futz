package service

import (
	"ahobson/gofutz/thing"
	"ahobson/gofutz/thing/mocks"
	"testing"
)

func TestServiceSimple(t *testing.T) {
	mockThing := &mocks.Thing{}
	expectedThingInput := thing.ThingInput{
		Id: 1,
	}
	// note how the mock is returning something that the actual
	// "thing" package could never return
	expectedThingOutput := thing.ThingOutput{Name: "Thing Name Wrong: 1"}
	mockThing.On("FetchThing", expectedThingInput).Return(expectedThingOutput)
	s := NewService(mockThing)
	r := s.DoTheThing(1)
	if r != expectedThingOutput.Name {
		t.Logf("Service return mismatch, expected %s, got %s",
			expectedThingOutput.Name, r)
		t.FailNow()
	}
}
