package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Vitor"
	var age = 20
	var version float32 = 1.1

	fmt.Println("Ola, Sr", name, "sua idade é", age)
	fmt.Println("Esta na versão ", version)

	fmt.Println("O tipo da variavel", name, "é", reflect.TypeOf(name))
	fmt.Println("O tipo da variavel", age, "é", reflect.TypeOf(age))
	fmt.Println("O tipo da variavel", version, "é", reflect.TypeOf(version))
}
