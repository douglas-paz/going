package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	saudacao()
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

func notaParaConceito(n float64) string {
	var nota = int(n)

	// Em Go é executado o primeiro case de sucesso e finalizado não precisando do break
	switch nota {
	case 10:
		fallthrough
		// sinaliza que caindo neste case é continuada a execução
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}
func saudacao() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite.")
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}
