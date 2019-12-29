package BehavioralType

import "testing"

func TestColleagueSeller_Colleguer(t *testing.T) {
	var (
		meitdor MeditorCompany
		seller *ColleagueSeller
		buyer *ColleagueBuyer
	)
	seller = &ColleagueSeller{meditor:meitdor}
	buyer = &ColleagueBuyer{meditor:meitdor}
	meitdor = &Meditor{
		name:   "58同城",
		buyer:  buyer,
		seller: seller,
	}
	// 卖家和卖家注册到中介
	seller.Colleguer(meitdor)
	buyer.Colleguer(meitdor)
	// 发布需求
	seller.Send("卖一套两室一厅100平米的Lofty")
	buyer.Send("求购一个两室一厅的房子")
}
