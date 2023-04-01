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

/*func CalculoConVectores(numeros... int,B... int) (int,int){
	var suma int
	var ProductoEscalar int
	for i:= 0; i < len(A); i++ {
		for j:= 0; j <= i; j++{
			suma = suma + (A[i]+B[j])
		}
	}

	return suma, ProductoEscalar
} */

//3. Escribir una función que reciba dos arreglos A y B, de N y M elementos respectivamente que representan conjuntos de enteros y
//devuelva una arreglo con la unión y otro con la intersección de A y B.
