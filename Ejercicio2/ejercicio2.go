package main

import "fmt"

func main(){
	n := calcularPromedio(1,2,3,4,5,6)
	fmt.Println(n)
}

func calcularPromedio(notas ...int)float64{

sumatoria := 0
for _, nota := range notas {
	sumatoria += int(nota)
}
promedio := float64(sumatoria)/ float64(len(notas))

return promedio
}