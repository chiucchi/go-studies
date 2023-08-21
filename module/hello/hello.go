// declara o package main
package main

// importa o pacote fmt e o pacote greetings que escrevemos previamente
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// requere uma mensagem de saudação
	// message, err := greetings.Hello("Giulia")

	// um slice de names.
	names := []string{"Giulia", "Fulano", "João"}

	// requere uma mensagem de saudação para cada nome
	messages, err := greetings.Hellos(names)

	// se um erro ocorrer, imprime o erro e termina o programa
	if err != nil {
		log.Fatal(err)
	}

	// se não houver erro, imprime as mensagens retornadas pela função Hello
	fmt.Println(messages)
}
