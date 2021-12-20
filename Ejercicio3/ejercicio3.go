package main

import "fmt"

func main(){
	s := calcularSalario(60, "A")
	fmt.Println(s)
}


func calcularSalario(minutos int, categoria string) float64{

	salario := 0.00
	var hora float64
	hora = float64(minutos)/60
	
	if categoria == "C"{
		salario = 1000 * float64(hora)
		
	}
	if categoria == "B" {
		salario = (1500 * hora) + (0.2 * (1500 * hora)) // 1800
		
	}

	if categoria == "A"{
		salario = ((3000 * hora) + (0.5 * (3000 * hora))) //4500
	}

	return float64(salario)
}