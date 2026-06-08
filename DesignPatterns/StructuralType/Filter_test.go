package StructuralType

import (
	"testing"
)

func TestHarmfulRubbishCriteria_RubbishFilter(t *testing.T) {
	rub := make([]Rubbish, 0)
	rub = append(rub, Rubbish{
		name:       "shell",
		isHarm:     false,
		isRecycled: false,
		isDry:      true,
		isWet:      false,
	})
	rub = append(rub, Rubbish{"ceramics", false, false, true, false})
	rub = append(rub, Rubbish{"vegetable waste", false, false, false, true})
	rub = append(rub, Rubbish{"fruit peel", false, false, false, true})
	rub = append(rub, Rubbish{"mercury thermometer", true, false, false, false})
	rub = append(rub, Rubbish{"battery", true, false, false, false})
	rub = append(rub, Rubbish{"bulb", true, false, false, false})
	rub = append(rub, Rubbish{"waste paper and plastic", false, true, false, false})
	rub = append(rub, Rubbish{"metal and cloth", false, true, false, false})
	rub = append(rub, Rubbish{"glass", false, true, false, false})

	dryFilter := DryRubbishCriteria{}
	wetFilter := WetRubbishCriteria{}
	harmFilter := HarmfulRubbishCriteria{}
	recyFilter := RecycledRubbishCriteria{}

	t.Log(dryFilter.RubbishFilter(rub))
	t.Log(wetFilter.RubbishFilter(rub))
	t.Log(harmFilter.RubbishFilter(rub))
	t.Log(recyFilter.RubbishFilter(rub))

}
