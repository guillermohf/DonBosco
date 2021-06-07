package main

import "fmt"

func main() {

    //Condional compuesta
	var num1, num2, num3 int

	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)


	if (num1 > num2 && num1 > num3){
		fmt.Print(num1)
	}else {
		if ( num2 > num3 ){
			fmt.Print(num2)
		}else{
			fmt.Print(num3)
		}
	}
}