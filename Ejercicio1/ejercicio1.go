package main 
import "fmt"

func main(){
	p := calcularDescuento(75000.00)
	fmt.Println(p)
}


func calcularDescuento(sueldo float64)float64{
	var sueldoDesc float64
	var descuento float64
	if sueldo > 50000.00 && sueldo < 150000.00 {
			descuento = 0.17
			sueldoDesc = sueldo * descuento
			return sueldoDesc
	} else{
		descuento = 0.27
		sueldoDesc = sueldo * descuento
		return sueldoDesc
	}
}