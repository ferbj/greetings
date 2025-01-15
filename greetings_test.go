package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Juan"
	//buscamos coincidencia exacta con esta regexp
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
