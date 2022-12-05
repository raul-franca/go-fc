package tax

import (
	"fmt"
	"testing"
)

func TestCalulateTax(t *testing.T) {
	amount := 500.0
	expect := 5.0

	result := CalulateTax(amount)

	if result != expect {
		fmt.Printf("O amount eh %f e o retorno foi %f", amount, result)
	}
}
func TestCalulateTaxBatch(t *testing.T) {

	type calcTax struct {
		amount, expect float64
	}
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalulateTax(item.amount)
		if result != item.expect {
			t.Errorf("O amount eh %f e o retorno foi %f", item.amount, result)
		}
	}

}

// go test -bench=. -run=ˆ#
func BenchmarkCalulateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalulateTax(500.0)
	}
}

func BenchmarkCalulateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalulateTax2(1000.0)
	}
}

// go test -fuzz=. -fuzztime=5s
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalulateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}
