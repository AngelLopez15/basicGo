package main

import "fmt"

func main() {
	// las estructuras en Go nos permiten almacenar una coleccion de campos.
	// similar a las clases en otros lenguajes de programacion

	// Para declarar una estructura se usa la palabra reservada type seguido del nombre de la estructura
	// Entre llaves se declarar cada uno de los campos que contendra la estructura y el tipo de dato
	type book struct {
		Name    string
		Author  string
		Country string
	}

	// Para crear instancias de la estructura debemos crear una variable, asignarle  el nombre de la estrutura
	// y dentro de las llaves poner el campo seguido del valor
	book1 := book{
		Name:    "A Carol Christmas",
		Author:  "Charles Dickens",
		Country: "England",
	}

	// Otra forma de crear instancias es atra vez de una estrutura literal
	book2 := book{
		"The adventures of the Sherlok Holmes",
		"Sir Artur Conan Doyle",
		"England",
	}

	// Podemos solo usar un campo de la estructura. Al no usar los otros campos
	// automaticamente se les asigna el valor cero
	// Sin embargo si solo solo se va a usar un campo no se puede hacer con intancias
	// literales
	book3 := book{
		Name: "Mi libro",
	}

	// Se pueden crear apuntadores a una estructura haciendo uso el operador de referenciacion
	miApuntador := &book1

	// Como estamos usando campos de una estructura si queremos cambiar el valor debemos poner
	// el operador de desrefenciacion entre parentesis seguido del operador punto para acceder al campo
	// (*miApuntador).Country = "United Kingdom"
	// Aunque Go reconoce que estamos accediendo al valor para reasignarlo atraves de un punturo
	// Entonces no es necesario usar la nomenclatura del parentesis y el asterisco
	miApuntador.Country = "United Kingdom"

	fmt.Printf("%+v\n", book1)
	fmt.Printf("%+v\n", book2)
	fmt.Printf("%+v\n", book3)

	//Si queremos acceder a un campo en especifico se hace mediante el operador punto
	fmt.Printf("%+v\n", book1.Author)
	fmt.Printf("Este mi apuntador -> %+v\n", miApuntador)
}
