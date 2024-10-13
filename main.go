package main

import (
	"fmt"
	"tigrinho/ascii"
	"tigrinho/slots"
)

func main() {
	var dineiros float64 = 1000.00
	var gastou float64 = 0.0
	var recebeu float64 = 0.0
	var count int = 0
	var ganhou int = 0
	var perdeu int = 0

	for {
		var input string

		fmt.Println("Escolha uma opcao:\n 1 - jogar \n 9 - para sair")

		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Tigrinho ganhou tudo, error")
		}

		if input == "9" {
			break
		}

		if input == "1" {

			count++

			if dineiros <= 3.5 {
				fmt.Println("Você não tem saldo suficiente! Gostaria da localização do agiota mais proximo para tentar a sorte novamente?")
				break
			}

			gastou += 3.5

			var resultado []string = slots.RodarSlots()
			fmt.Println(resultado)

			if slots.ChecarResultado() {
				dineiros += 10.5
				ganhou++
				recebeu += 3.5
				fmt.Println("TIGRINHO ACREANO TA PAGANDO MUUUUITO")
				fmt.Println(ascii.Ganhou())

			} else {
				dineiros -= 3.5
				perdeu++
				fmt.Println("Não vai ter carta agora")
				fmt.Println(ascii.Perdeu())
			}

			fmt.Println("Voce já jogou", count, "vezes")
			fmt.Println("Já ganhou um total de ", ganhou, "vezes")
			fmt.Println("Perdeu", perdeu, "vezes kkkkkkkkkkkkkkkkkk")
			fmt.Println("O seu saldo atual é de: ", dineiros)
			fmt.Println("O seu NET PROFIT é de:", recebeu-gastou)
		}

		fmt.Println("Comando invalido, tente novamente")
	}
}
