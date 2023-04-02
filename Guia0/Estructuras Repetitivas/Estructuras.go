/* Estructuras repetitivas*/

package estructuras

import (
	"fmt"
)

//1. Escribir una función que recibe un número entero n, no negativo, y devuelve su factorial.*/

func Factorial(numero int) {
	factorial := numero

	if numero > 0 {
		for i := 1; i < numero; i++ {
			factorial = factorial * i
		}
		fmt.Printf("El factorial de %d! es: %d", numero, factorial)
	}
	if numero <= 0 {
		fmt.Printf("El numero debe ser entero y positivo")
	}
}

// 2. Escribir una función que recibe dos enteros: a y b y devuelve el producto. Para el cálculo, utilizar sumas sucesivas.
func CalcularProducto(a, b int) {
	producto := 0
	for i := 1; i <= b; i++ {
		producto = producto + a
	}
	fmt.Println(producto)
}

//3. Escribir una función que indique si un número entero, ingresado por un usuario, es primo (devuelve verdadero o falso).

func ValidarSiEsPrimo(numero int) {
	EsPrimo := true
	if numero == 0 || numero == 1 {
		fmt.Printf(" %d no es un numero primo\n", numero)
	} else {
		for i := 2; i <= numero/2; i++ {
			if numero%i == 0 {
				fmt.Printf(" %d no es un numero primo\n", numero)
				EsPrimo = false
				break
			}
		}
		if EsPrimo == true {
			fmt.Printf(" %d es un numero primo\n", numero)
		}
	}
}
