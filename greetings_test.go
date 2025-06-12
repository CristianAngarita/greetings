package greetings

import (
	"regexp"
	"testing"
)

func TestHelloNAme(t *testing.T) {
	name := "Steven"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Steven")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Steven) = %q, %v , quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v , quiere "" , error`, msg, err)
	}
}
