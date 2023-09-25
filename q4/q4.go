package q4

import "fmt"

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	var x int
	if len(prices) == 1 {
		x = 3
	}
	if len(prices) == 0 {
		return 0, fmt.Errorf("Valor inválido")
	}
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			if x == 2 {
				x = 3
			} else if x == 0 || x == 1 {
				x = 1
			}
		} else if prices[i+1] < prices[i] {
			if x == 1 {
				x = 3
			} else if x == 0 || x == 1 {
				x = 2
			}
		}

	}

	return x, nil
}
