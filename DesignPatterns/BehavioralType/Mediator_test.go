package BehavioralType

import "testing"

func TestColleagueSeller_Colleguer(t *testing.T) {
	var (
		meitdor MeditorCompany
		seller  *ColleagueSeller
		buyer   *ColleagueBuyer
	)
	seller = &ColleagueSeller{meditor: meitdor}
	buyer = &ColleagueBuyer{meditor: meitdor}
	meitdor = &Meditor{
		name:   "HousingAgency",
		buyer:  buyer,
		seller: seller,
	}

	seller.Colleguer(meitdor)
	buyer.Colleguer(meitdor)

	seller.Send("sell a 100 sqm two-bedroom loft")
	buyer.Send("want to buy a two-bedroom apartment")
}
