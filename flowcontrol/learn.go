package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func Learn() {
	fmt.Println("Learn Flow Control")

	// if else
	score := 50
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}

	// switch case

	switch score {
	case 80:
		fmt.Println("A")
	case 70:
		fmt.Println("B")
	case 60:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	// switch case with conditions
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other OS")
	}

	// Read file
	file, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(file))
	}

	// for loop
	for i := 10; i >= 0; i-- {
		fmt.Println("Time out : ", i)
		if i == 0 {
			fmt.Println("Boom!")
			break
		}
		time.Sleep(time.Second)
	}
}
