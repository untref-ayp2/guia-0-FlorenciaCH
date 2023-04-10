package menu

import "fmt"

/*2. Formar un menú con 4 opciones (como se muestra debajo) y al elegir una de ellas
 mostrar un cartel diciendo qué opción se eligió o si fue una opción incorrecta.
- Opción 1
- Opción 2
- Opción 3
- Opción 4*/

func MostrarMenu() {

	fmt.Print("Elija una opción del siguiente menú:\n", "Opción 1\n", "Opción 2\n", "Opción 3\n", "Opción 4\n")

}

func ElegirOpcion(opcion int) {
	switch opcion {
	case 1:
		println("Usted ha elegido la Opción 1")
	case 2:
		println("Usted ha elegido la Opción 2")
	case 3:
		println("Usted ha elegido la Opción 3")
	case 4:
		println("Usted ha elegido la Opción 4")

	default:
		fmt.Println("La opcion elegida es incorrecta")
	}

}
