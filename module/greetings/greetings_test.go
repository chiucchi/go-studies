// qualquer file terminando em _test.go é reconhecido como um arquivo de teste

package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName checa um retorno válido passando um nome para a função Hello
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty chama a função Hello com um argumento vazio, validando o retorno de um erro
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
