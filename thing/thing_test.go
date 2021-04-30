package thing

import (
	"testing"
)

func TestThingSimple(t *testing.T) {
	rt := &RealThing{}
	to, err := rt.FetchThing(ThingInput{})
	if err != nil {
		t.Logf("Unexpected Error %s", err)
		t.FailNow()
	}
	expectedName := "Thing Name: 0"
	if (to.Name != expectedName) {
		t.Logf("Name mismatch, expected %s, got %s", expectedName, to.Name)
		t.FailNow()
	}
}

func TestThingError(t *testing.T) {
	rt := &RealThing{}
	_, err := rt.FetchThing(ThingInput{Id: -1})
	if err == nil {
		t.Logf("No Error on negative ID!")
		t.FailNow()
	}
}
