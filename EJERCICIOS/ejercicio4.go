package main

import "fmt"

func main() {

    //Condional compuesta
	var num1, num2, num3, promedio int

	fmt.Scan(num1)
	fmt.Scan(num2)
	fmt.Scan(num3)

	promedio = ((num1+num2+num3)/3)




	if (promedio >= 7){
		fmt.Print("Promocionado")
	}else{
		if (promedio >= 4){
			fmt.Print("Regular")
		}else
		{
			fmt.Print("Regular")
		}
	}
    


}