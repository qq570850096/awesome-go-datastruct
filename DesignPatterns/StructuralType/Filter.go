package StructuralType

type Rubbish struct {
	name       string
	isHarm     bool
	isRecycled bool
	isDry      bool
	isWet      bool
}

type Criteria interface {
	RubbishFilter(rubbishs []Rubbish) []Rubbish
}

type DryRubbishCriteria struct{}

func (DryRubbishCriteria) RubbishFilter(rubbishs []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishs {
		if v.isDry == true {
			res = append(res, v)
		}
	}
	return res
}

type WetRubbishCriteria struct{}

func (WetRubbishCriteria) RubbishFilter(rubbishs []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishs {
		if v.isWet == true {
			res = append(res, v)
		}
	}
	return res
}

type HarmfulRubbishCriteria struct{}

func (HarmfulRubbishCriteria) RubbishFilter(rubbishs []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishs {
		if v.isHarm == true {
			res = append(res, v)
		}
	}
	return res
}

type RecycledRubbishCriteria struct{}

func (RecycledRubbishCriteria) RubbishFilter(rubbishs []Rubbish) []Rubbish {
	res := make([]Rubbish, 0)
	for _, v := range rubbishs {
		if v.isRecycled == true {
			res = append(res, v)
		}
	}
	return res
}
