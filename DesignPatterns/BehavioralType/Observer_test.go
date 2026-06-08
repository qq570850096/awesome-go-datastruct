package BehavioralType

import "testing"

func TestPlatform_Attach(t *testing.T) {

	platform := Platform{list: []IReader{}}

	reader := Reader{name: "A"}

	platform.Attach(&reader)

	reader2 := Reader{name: "B"}

	platform.Attach(&reader2)
	platform.Change("Go Core Programming")

	platform.Detach(&reader2)
	platform.Change("Advanced Go Programming")
}
