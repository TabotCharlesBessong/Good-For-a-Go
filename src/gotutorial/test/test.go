package testpackage

import "fmt"


func MyFunction(step int) {
	if step <=7 {
		fmt.Println("How are you doing?")

	}else if step >= 8 && step < 16{

		fmt.Println("Am fine and you?")
	}else{

		fmt.Println("Am also doing great")
	}
}
