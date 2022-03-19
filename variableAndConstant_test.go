package main

import (
	"testing"
)

func TestVariabesConstant(t *testing.T) {
	got := BMICalculator()
	res := 0.0018365472910927456
	if got != res {
		t.Errorf("BMICalculator() = %f; want %f  ", got, res)
	}

}
