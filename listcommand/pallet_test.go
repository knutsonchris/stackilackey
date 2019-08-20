package listcommand

import "testing"

func TestPallet_Pallet(t *testing.T) {
	// run a list pallet command and make sure the default Stacki pallet exists
	p := Pallet{}
	pallets, err := p.Pallet("stacki", "", "", "", "")
	if err != nil {
		t.Fatalf("list pallet stacki failed with error %s", err)
	}
	for _, pallet := range pallets {
		if pallet.Name == "stacki" {
			return
		}
	}
	t.Error("list pallet stacki failed. unable to find expected stacki pallet")
}
