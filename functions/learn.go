package functions

import "fmt"

var Exportvar string = "Export variable"

func Learn() {
	fmt.Println("Learn functions")
	fmt.Println("1 + 2 = ", add(1, 2))
	fmt.Println(convertor(10))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

// func name(parameters type) return type
func add(x, y int) int {
	return x + y
}

func convertor(x float64) (float64, float64) {
	f := x * 0.3048
	i := x * 0.0254
	return f, i
}

func sum(x ...int) int {
	total := 0
	for index, value := range x {
		fmt.Println("index: ", index, "value: ", value)
		total += value
	}
	return total

}
