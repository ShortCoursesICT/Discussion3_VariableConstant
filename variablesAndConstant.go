package main

import "fmt"

const (
	Weight = 50
	Height = 165
	Age    = 22
)

func BMICalculator() float64 {
	//Insert your code here
	BMI := Weight / (Height * Height)
	//print out the result, it will be zero. We need to make changes here for it to work.
	fmt.Println(BMI)
	//Do not remove this line
	return float64(BMI)

}

func main() {
	BMICalculator()
}
