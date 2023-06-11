package pointers

import "fmt"

/*
& วางไว้หน้าคือ การอ้างอิงถึง address ของตัวแปร
* วางไว้หน้าคือ การอ้างอิงถึง value ของตัวแปร
* วางไว้หน้า type คือ การประกาศว่าเป็น pointer เช่น *int, *string, *struct (เก็บค่า memory address ของตัวแปร ว่างๆหรือ nil)
*/
func Learn() {
	fmt.Println("Learn Pointers")

	// Print the memory address of x
	x := 10
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println("--------------------")
	y := x
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Println("--------------------")
	// Declare a pointer variable
	z := &x
	fmt.Println(z)
	fmt.Println("Read value from Pointer z", *z) //dereferencing

	fmt.Println("--------------------")
	b := student{name: "Bank"}
	setName(&b)
	fmt.Println(b)
}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "John"
}
