package main

import (
	"github.com/Andrew-Wichmann/FundraiserSiteAPI/pkg/orm"
)

func main() {
	pledge := orm.Pledge{Email: "awichmann@mintel.com", PlanID: "2", Rate: 0.1, Maximum: 20}
	err := pledge.Save()
	if err != nil {
		panic(err)
	}
}
