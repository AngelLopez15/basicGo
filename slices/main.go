package main

import "fmt"

func main() {
	// Slices
	// los slices nos permiten trabajar con arrays de manera dinamica.
	// Son apuntadores a arrays (no almacenan ningún dato), permiten trabajar con arrays de forma dinámica.
	set := [7]string{"🥩", "🍖", "🍗", "🍕", "🙉", "🙈", "🥵"}

	// creando el primer slice, debemos indicar cual sera el arrary al que vamos a apuntar
	// entre corchetes poner el indice inicial "ini" dos puntos y el indice final "final" al que vamos a apuntar
	// caras := set[ini:final]
	// IMPORTANTE El ultimo elemento del rango del slice es excluyente. es decir no se incluye
	caras := set[4:7]
	comida := set[0:4]

	// Si modificamos el valor de un arreglo en un slice, este va a cambiar su valor en todos los arreglas de donde venga
	// es decir Modificar elementos del slice, en verdad es modificar elementos del array donde apunta el slice.
	// los slices no poseen datos simplemente apuntan al array
	comida[3] = "🍆"
	fmt.Println("Array original:", set)
	fmt.Println("Caras:", caras)
	fmt.Println("Comida:", comida)
	fmt.Println("Comida:", comida[3])

	// Si en un arreglo no se especifica el valor inicial o el valor final, GO lo interpretara como que debe tomar desde el primer
	// valor del arreglo o hasta el ultimo valor del arreglo respectivamente
	// caras := set[:7]
	// comida := set[0:]
	// no se especifica ninguno de los dos entonces imprimira todos
	// comida := set[:]
}
