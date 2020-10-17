package golangtest

import (
	"strconv"
)

// Return multiple values
func MultipleValues() (int, int) {
	// variable declaration
	var firstNum int
	//variable definition
	firstNum = 3
	secondNum := 5
	return firstNum, secondNum
}

// Pointer to pointer
func PointerToPointer() **int {
	var v int = 100
	var pt1 *int = &v
	var pt2 **int = &pt1
	return pt2
}

// Maps
func MarksMap(maths, english int) map[string]int {
	marks := make(map[string]int)
	marks["Maths"] = 100
	marks["English"] = 95
	return marks
}

// Closures
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Type Casting
func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func FloatToInt(f float64) int {
	return int(f)
}

func IntToFloat(i int) float64 {
	return float64(i)
}

// Pass by pointer
func IncrementOne(i *int) {
	*i+=1
}

func main() {

}
