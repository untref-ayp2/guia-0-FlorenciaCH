//ARREGLOS//

package arreglos

//1. Escribir una función que reciba un arreglo de enteros como parámetros y devuelva la suma de todos sus elementos

func SumarEnteros(numeros ...int) {
	sumatoria := 0
	for i := 0; i < len(numeros); i++ {
		sumatoria = sumatoria + numeros[i]
	}
	println(sumatoria)
}

//2. Escribir una función que reciba dos arreglos (que representan vectores) del mismo tamaño
//y devuelvan la suma y el producto escalar de los vectores

func CalculoConVectores(A []int, B []int) (suma []int, producto int) {
	suma = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		suma[i] = A[i] + B[i]
		producto = producto + A[i]*B[i]
	}
	return suma, producto
}

//3. Escribir una función que reciba dos arreglos A y B, de N y M elementos respectivamente que representan conjuntos de enteros y
//devuelva una arreglo con la unión y otro con la intersección de A y B.

func OperaciónConArreglos(A []int, B []int) (union []int, interseccion []int) {

	union = append(A, B...)

	for _, primerElementoA := range A {
		for _, primerElementoB := range B {
			if primerElementoA == primerElementoB {
				interseccion = append(interseccion, primerElementoA)
			}
		}
	}
	return union, interseccion
}
