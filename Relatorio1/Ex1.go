package main

import( 
	"fmt"
)

func ParOuImpar(n int) (string) {
	if n % 2 == 0 {
		return "Par"
	} else {
		return "Impar"
	}
}

func main() {
	fmt.Println(ParOuImpar(9));
}