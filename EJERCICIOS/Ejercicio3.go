package main

import "fmt"

func main() {

    //Condional compuesta
	var num1, num2 int

	fmt.Scan(num1)
	fmt.Scan(num2)

	fmt.Print(num1 > num2 ? num1 : num2 )

	if (num1 > num2){

		fmt.Print(num1,"es mayo que",num2)
	}else{
		fmt.Print(num2,"es mayor que",num1)
	}
    
	


}