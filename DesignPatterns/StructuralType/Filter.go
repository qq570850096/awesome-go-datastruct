package StructuralType

type Rubbish struct {
	name string
	isHarm bool
	isRecycled bool
	isDry bool
	isWet bool
}

// 我们过滤的标准接口，即一个抽象的过滤器
type Criteria interface {
	// 定义过滤的标准
	RubbishFilter(rubbishs []Rubbish) []Rubbish
}

// 具体的过滤类
// 干垃圾
type DryRubbishCriteria struct {}

func (DryRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
	res := make([]Rubbish,0)
	for _,v := range rubbishs {
		if v.isDry == true {
			res = append(res,v)
		}
	}
	return res
}

// 湿垃圾
type WetRubbishCriteria struct {}

func (WetRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
	res := make([]Rubbish,0)
	for _,v := range rubbishs {
		if v.isWet == true {
			res = append(res,v)
		}
	}
	return res
}

// 有害垃圾
type HarmfulRubbishCriteria struct {}

func (HarmfulRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
	res := make([]Rubbish,0)
	for _,v := range rubbishs {
		if v.isHarm == true {
			res = append(res,v)
		}
	}
	return res
}

// 可回收垃圾
type RecycledRubbishCriteria struct {}

func (RecycledRubbishCriteria)RubbishFilter(rubbishs []Rubbish) []Rubbish  {
	res := make([]Rubbish,0)
	for _,v := range rubbishs {
		if v.isRecycled == true {
			res = append(res,v)
		}
	}
	return res
}


