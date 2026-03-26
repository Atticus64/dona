package tests

import (
	"fmt"
	"regexp"
	"testing"
)

func Saludar(nombre string) string {
	saludo := fmt.Sprintf("Hola %s", nombre)
	return saludo
}

func TestHelloName(t *testing.T) {
    name := "Juan Carlos Bodoque"
    want := regexp.MustCompile(`\b`+name+`\b`)
	// Must be wrong 
    //msg := Saludar("Pedro Paramo")
	// Must be good 
    msg := Saludar("Juan Carlos Bodoque")
    if !want.MatchString(msg) {
        t.Errorf(`Saludar("Juan Carlos Bodoque") = %q, want match for %#q, nil`, msg, want)
    }
}
