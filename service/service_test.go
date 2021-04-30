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
	mockThing.On("FetchThing", expectedThingInput).Return(expectedThingOutput, nil)
	s := NewService(mockThing)
	r, err := s.DoTheThing(1)
	if err != nil {
		t.Logf("Unexpected Error: %v", err)
		t.FailNow()
	}
	if r != expectedThingOutput.Name {
		t.Logf("Service return mismatch, expected %s, got %s",
			expectedThingOutput.Name, r)
		t.FailNow()
	}
}

func TestServiceNoFail(t *testing.T) {
	mockThing := &mocks.Thing{}
	expectedThingInput := thing.ThingInput{
		Id: -1,
	}
	// note how the mock does not return an error on negative id
	expectedThingOutput := thing.ThingOutput{Name: "Thing Name Wrong: 1"}
	mockThing.On("FetchThing", expectedThingInput).Return(expectedThingOutput, nil)
	s := NewService(mockThing)
	r, err := s.DoTheThing(-1)
	if err != nil {
		t.Logf("Unexpected Error: %v", err)
		t.FailNow()
	}
	if r != expectedThingOutput.Name {
		t.Logf("Service return mismatch, expected %s, got %s",
			expectedThingOutput.Name, r)
		t.FailNow()
	}
}
