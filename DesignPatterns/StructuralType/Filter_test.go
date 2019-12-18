package StructuralType

import (
	"testing"
)

func TestHarmfulRubbishCriteria_RubbishFilter(t *testing.T) {
	rub := make([]Rubbish,0)
	rub = append(rub,Rubbish{
		name: 		"果壳",
		isHarm:     false,
		isRecycled: false,
		isDry:      true,
		isWet:      false,
	})
	rub = append(rub,Rubbish{"陶瓷",false,false,true,false})
	rub = append(rub,Rubbish{"菜根菜叶",false,false,false,true})
	rub = append(rub,Rubbish{"果皮",false,false,false,true})
	rub = append(rub,Rubbish{"水银温度计",true,false,false,false})
	rub = append(rub,Rubbish{"电池",true,false,false,false})
	rub = append(rub,Rubbish{"灯泡",true,false,false,false})
	rub = append(rub,Rubbish{"废纸塑料",false,true,false,false})
	rub = append(rub,Rubbish{"金属和布料",false,true,false,false})
	rub = append(rub,Rubbish{"玻璃",false,true,false,false})

	dryFilter := DryRubbishCriteria{}
	wetFilter := WetRubbishCriteria{}
	harmFilter := HarmfulRubbishCriteria{}
	recyFilter := RecycledRubbishCriteria{}

	t.Log(dryFilter.RubbishFilter(rub))
	t.Log(wetFilter.RubbishFilter(rub))
	t.Log(harmFilter.RubbishFilter(rub))
	t.Log(recyFilter.RubbishFilter(rub))

}
