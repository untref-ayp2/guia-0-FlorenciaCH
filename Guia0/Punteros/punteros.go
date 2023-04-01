package punteros

//<h4>Punteros</h4>

//1. Escribir una función con la siguiente firma func swap(px, py *int) que permita
//intercambiar el valor almacenado en dos variables enteras a través de punteros.

func swap(px, py *int) (int, int) {
	*px = *py
	*py = *px
	return *py, *px
}
