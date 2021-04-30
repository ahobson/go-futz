package thing

import (
	"testing"
)

func TestThingSimple(t *testing.T) {
	rt := &RealThing{}
	to := rt.FetchThing(ThingInput{})
	expectedName := "Thing Name: 0"
	if (to.Name != expectedName) {
		t.Logf("Name mismatch, expected %s, got %s", expectedName, to.Name)
		t.FailNow()
	}
}
