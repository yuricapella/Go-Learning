package main

func sintaxeArrayValoresZero() {
	var arrayString [5]string
	arrayString[0] = "Posição 1"

	var arrayInt [5]int
	arrayInt[0] = 10
}

func sintaxeArrayInicializacao() {
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	_ = array2
}

func sintaxeArrayTamanhoAutomatico() {
	array3 := [...]int{1, 2, 3, 4, 5}
	_ = array3
}

func sintaxeSliceInicializacao() {
	slice := []int{1, 2, 3, 4, 5}
	_ = slice
}

func sintaxeSliceAppend() {
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	slice = append(slice, 7, 8, 9)
}

func sintaxeSliceMake() {
	slice2 := make([]int, 5)
	slice2ComCap := make([]int, 3, 10)
	_ = slice2
	_ = slice2ComCap
}

func sintaxeSliceSlicing() {
	array3 := [...]int{1, 2, 3, 4, 5}
	slice3 := array3[1:3]
	array3[1] = 20
	_ = slice3
}
