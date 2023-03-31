package Polinomio

import (
	"fmt"
)

/* 1. Definir una función que, dado los coeficientes de un polinomio de grado n (números flotantes)
muestre por pantalla el polinomio completo,
por ejemplo si recibe los coeficientes 3.0, 2.0 y 1.0 debe mostrar 3.0 + 2.0 X + 1.0 X**2 */

func ImprimirPolinomio(coeficientes ...float32) {

	for i := 0; i < len(coeficientes); i++ {
		if i > 1 {
			fmt.Printf("%+.1f X**%v", coeficientes[i], i) // (%+ .1f) toma solo un decimal.
		} else if i == 1 {
			fmt.Printf("%+.1f X", coeficientes[i])
		} else {
			fmt.Printf("%.1f", coeficientes[i])
		}
	}
}
