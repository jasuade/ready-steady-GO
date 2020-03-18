package main

import "fmt"

type If interface {
	MetodoGenerico()
}

type Estructura struct {
	Word string
}

func (estruct *Estructura) MetodoGenerico() {
	if estruct == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(estruct.Word)
}

func main() {
	var interfazEjemplo If
	var estructuraEjemplo *Estructura //Al no asignarle valor obtenemos nil
	interfazEjemplo = estructuraEjemplo
	describe(interfazEjemplo)
	interfazEjemplo.MetodoGenerico()

	interfazEjemplo = &Estructura{"Random"}
	describe(interfazEjemplo)
	interfazEjemplo.MetodoGenerico()
}

func describe(interfazEjemplo If) {
	fmt.Printf("(%v, %T)\n", interfazEjemplo, interfazEjemplo)
}
