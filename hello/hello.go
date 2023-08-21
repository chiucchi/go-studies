// declara um package main
package main

// importa o pacote fmt - uma lib standard do Go
// o segundo import é de um modulo externo, para baixar ele só digitar no terminal 'go mod tidy'
import (
	"fmt"

	"rsc.io/quote"
)

// declara uma função main, que executa o programa por padrão ao rodar o package main
func main() {
	fmt.Println("Hello, World!") // printa hello world usando o pacote fmt
	fmt.Println(quote.Go())      // printa uma frase usando o pacote rsc.io/quote
	fmt.Println(quote.Opt())     // printa usando outra função presente no pacote rsc.io/quote
}
