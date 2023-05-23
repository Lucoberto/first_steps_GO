package main

/* fmt se usa para dar formato, acepta carecteres especiales  */
import (
	"fmt"
)

/* Funcion en GO */
func main() {
	/* Imprimir y concatenar */
	fmt.Println("Its no my name" + "but yes")

	/* Hay que de clarar variables */
	var hello = "wellcome"
	fmt.Println(hello)

	/* Puedes evitar declarar variables usando :=
	pero esto solo dentro de funciones. */
	wellcome := "My name is"
	fmt.Printf(wellcome + hello)

	/* Bucle con for */
	numero := 1
	for numero <= 100 {
		fmt.Println(numero)
		numero = numero + numero
	}
	/* Bucle for comprimido */
	for numerito := 1; numerito <= 100; numerito++ {
		fmt.Println(numerito)
	}

	/* If funciona como en python */

	if hello == "wellcome" {
		fmt.Println("works!!")
	} else {
		fmt.Println("Don't work!!")
	}

	/* Declaracion de variable dentro del if */
	if num := 200; num < 0 {
		fmt.Println(num, "este es tu numero")
	} else if num == 200 {
		fmt.Println(num, "este es tu numero")
	}

	/* Lo mismo que un if pero mas corto */
	nums := 200
	switch nums {
	case 100:
		fmt.Println("case 100")
	case 200:
		fmt.Println("case 200")
	default:
		fmt.Println("WoW Te pasas")
	}

	/* Se le da una variable a la funcion y un parametro para despues declarar el tipo en este caso interface,
	Luego se le da un valor al parametro de la funcion que se comprueba cada vez hasta haber comprobado
	todas las posibilidades. myfunc(true) !!! ES UNA FUNCION ANONIMA !!!
	*/
	myfunc := func(parameter interface{}) {
		switch t := parameter.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	myfunc(true)
	myfunc(1)
	myfunc("hey")
}
