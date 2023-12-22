package items

import "testing"

func TestAllProductsAreReturned(t *testing.T) {
	products := GetProducts()
	numberOfProducts := len(*products)
	if numberOfProducts != 4 {
		t.Fail()
		t.Log("Expected 4 products but found", numberOfProducts)
	}
}

