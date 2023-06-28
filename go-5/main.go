package main

import (
	M "pack/api-gin/models"
	R "pack/api-gin/routes"
)

func main() {
	M.Alunos = []M.Aluno{
		{Nome: "Vitor", CPF: "90090090088", RG: "89097097x"},
		{Nome: "Sara", CPF: "99990090088", RG: "89097097a"},
	}
	R.HandleRequests()
}
