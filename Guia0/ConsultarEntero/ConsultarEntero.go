package consultarentero

import (
	"fmt"
)

//3. Escribir una función que reciba una lista de enteros y devuelva el menor y el mayor de la lista//

func ObtenerMayoryMenor(numeros ...int) {

	valorMinimo := numeros[0]
	valorMaximo := numeros[0]

	for i := 1; i < len(numeros); i++ {

		if numeros[i] < valorMinimo {
			valorMinimo = numeros[i]
		}
		if numeros[i] > valorMaximo {
			valorMaximo = numeros[i]
		}
	}

	fmt.Printf("Su valor minimo es: %d\n", valorMinimo)
	fmt.Printf("Su valor máximo es: %d\n", valorMaximo)
}
