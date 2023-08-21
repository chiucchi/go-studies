package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello retorna uma mensagem de saudação que incorpora o nome da pessoa recebida.
// o primeiro paratenses representa o parametro esperado pela função e seu tipo
// o segundo paratenses representa o tipo de retorno da função, separado por vírgula e entre parenteses se for mais de um valor
func Hello(name string) (string, error) {
	// se nenhum nome foi informado, retorna um erro
	if name == "" {
		return "", errors.New("empty name")
	}

	// um shortchut para criar uma variável e atribuir um valor
	// retorna a mensagem se recebeu um nome
	message := fmt.Sprintf(randomFormat(), name)

	// aqui o nil é usado para indicar que não houve erro
	return message, nil
}

// essa função recebe um slice de strings (nomes) e retorna um map de strings (mensagens)
func Hellos(names []string) (map[string]string, error) {
	// inicializa um map para armazenar as mensagens de saudação
	messages := make(map[string]string)
	// através de um laço de repetição chama a função de saudação para cada nome do slice
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// associa cada mensagem com o nome correspondente no map
		messages[name] = message
	}
	return messages, nil
}

// essa função retorna uma mensagem aleatória de boas-vindas
func randomFormat() string {
	// é como se fosse um array de strings
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Ciao, %v. Benvenuto!",
	}

	// retorna uma mensagem aleatória ao passar um valor de índice aleatório para o array
	return formats[rand.Intn(len(formats))]
}
