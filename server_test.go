package server

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Ethan"
	want := regexp.MustCompile(`Hello Ethan from SimpleHelloServer!`)
	server := SimpleHelloServer{ServerName: "SimpleHelloServer"}
	msg := server.Greet(name)
	if !want.MatchString(msg) {
		t.Errorf(`Hello("Ethan") = %q, want match for %#q, nil`, msg, want)
	}
}
