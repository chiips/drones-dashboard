package models

import "testing"

//init test values
var (
	lat1 = 83.1424
	lat2 = 42.0967
	lon1 = -52.8805
	lon2 = -141.2985
)

func TestHsin(t *testing.T) {

	//set up a couple tests to run through
	type test struct {
		data   float64
		answer float64
	}
	tests := []test{
		test{lat2 - lat1, 0.9895309499417008},
		test{lon2 - lon1, 0.05051973936121625},
	}

	for _, v := range tests {
		f := hsin(v.data)
		//compare result to answer
		if f != v.answer {
			t.Error("got", f, "want", v.answer)
		}
	}

}

func BenchmarkHsin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hsin(lat2 - lat1)
	}
}

func TestDistance(t *testing.T) {

	f := distance(lat1, lat2, lon1, lon2)

	answer := 5.352710712389778e+06

	if f != answer {
		t.Error("got", f, "want", answer)
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distance(lat1, lat2, lon1, lon2)
	}
}
