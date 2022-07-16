package orm

import (
	"testing"

	"github.com/asaskevich/govalidator"
)

func TestPledge(t *testing.T) {
	tests := []struct {
		name   string
		valid  bool
		pledge Pledge
	}{
		{name: "Valid Pledge", valid: true, pledge: Pledge{Email: "test@gmail.com", PlanID: "0", Maximum: 1, Rate: .1}},
		{name: "Float Maximum", valid: true, pledge: Pledge{Email: "test@gmail.com", PlanID: "0", Maximum: 11.1, Rate: .1}},
		{name: "Bad email", valid: false, pledge: Pledge{Email: "test", PlanID: "0", Maximum: 1, Rate: .1}},
		{name: "No email", valid: false, pledge: Pledge{PlanID: "0", Maximum: 1, Rate: .1}},
		{name: "No planID", valid: false, pledge: Pledge{Email: "test@gmail.com", Maximum: 1, Rate: .1}},
		{name: "No maximum", valid: false, pledge: Pledge{Email: "test@gmail.com", PlanID: "0", Rate: .1}},
		{name: "No rate", valid: false, pledge: Pledge{Email: "test@gmail.com", PlanID: "0", Maximum: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, _ := govalidator.ValidateStruct(tt.pledge)
			if tt.valid != valid {
				t.Logf("Expected: %t, received: %t", tt.valid, valid)
				t.Fail()
			}
		})
	}
}
