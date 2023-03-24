package main

import (
	"strconv"
	"fmt"
)

/* 1. Definir una función que, dado los coeficientes de un polinomio de grado n (números flotantes)
 muestre por pantalla el polinomio completo,
 por ejemplo si recibe los coeficientes 3.0, 2.0 y 1.0 debe mostrar 3.0 + 2.0 X + 1.0 X**2
2. Formar un menú con 4 opciones (como se muestra debajo) y
al elegir una de ellas mostrar un cartel diciendo qué opción se eligió o si fue una opción incorrecta.
- Opción 1
- Opción 2
- Opción 3
- Opción 4

3. Escribir una función que reciba una lista de enteros y devuelva el menor y el mayor de la lista */

func ImprimirPolinomio(coeficientes ...float32) {

	for i := 0; i < len(coeficientes); i++ {

		if i > 1 {
			coeficientes[i] = fmt.Sprintf(" %+.1f X**%v ", int(coeficientes[i]), i) // (%+ .1f) Con este ejemplo el input sería "-3.1416" y el output sería "- 3.1 ""
		} else if i == 1 {
			coeficientes[i] = fmt.Sprintf(" %+.1f X ", int(coeficientes[i]))
		} else {
			coeficientes[i] = fmt.Sprintf(" %.1f ", int(coeficientes[i]))
		}
	}

	for _, coeficiente := range arreglo {
		fmt.Print(coeficiente)
	}
}
