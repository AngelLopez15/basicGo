package main

import "fmt"

func main() {
	// Para crear un mapa en Go debemos hacer uso de la funcion make
	// dentro usar la palabra map y debemos especificar el tipo de valores que contendra
	animals := make(map[string]string)
	// Para asignar valores al mapa que se ha creado debemos poner el nombre del mapa
	// entre corchetes el nombre de la llave y asignarle el valor
	animals["gato"] = "🐱"
	animals["perro"] = "🐶"

	// Creando un mapa literal

	// Siempre debemos poner la como al final de cada uno de los elementos
	// hasta en el ultimo elemento debe llevar como obligatoriamente en GO
	fruits := map[string]string{
		"apple": "🍎",
		"pinia": "🍍",
	}

	// Para eliminar elementos de un mapa debemos usar la funcion delete
	// como primer argumento debe ir el nombre del mapa y como segundo el nombre de
	// la llave a eliminar.
	// imprimimos antes del delete para poder ver que efectivamente despues elimina el elemento
	fmt.Println("mapa frutas: ", fruits)
	delete(fruits, "pinia")

	fmt.Println("mapa animales: ", animals)
	fmt.Println("mapa frutas: ", fruits)

	// En los mapas podemos obtener sus elementos
	// para eso debemos poner el nombre de la mapa seguido de corchetes que dentro debe tener
	// el nombre de la llave al que queremos acceder
	fmt.Println("obteniendo al gato: ", animals["gato"])

	// Que sucede si queremos recuperar un elemento que no hemos creado
	// Cuando no existe go nos retorna el tipo de dato del valor que contiene el mapa
	// en este caso como almacena string regresaria un string vacio.
	// Cuando se recuperan valores a traves de la forma -> animals["gato"] Go nos puede
	// retornar dos valores (Go tiene la caracteristica de que puede retornar multiples valores)
	// Cuando se consulta un valor de esta manera -> animals["gato"] siempre  nos retorna dos valores
	// El primer valor -> Es el contenido
	// El segundo valor -> Es si la llave existe o no (un booleano) por convencion esta se llama "ok" pero puede ser cualquier nombre
	emoji, ok := animals["gato"]
	fmt.Println(emoji, ok)

	// Entonces si una llave no existe podemos hacer una validacion para que la cree
	// Cuando no usamos una variable el compilar da error, en este caso como no ocupamos el
	// primer parametro debemos sustituirlo por el blanck ( _ ) <- guion bajo
	if _, ok2 := animals["gorila"]; !ok2 {
		animals["gorila"] = "🦍"
	}
	fmt.Println("nuevo animals", animals)
}
