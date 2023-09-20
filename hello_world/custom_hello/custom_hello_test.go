package custom_hello

import (
	"regexp"
	"testing"
)

func TestCustomHello(t *testing.T) {

	// Testing for a valid return type:
	name := "Aymane"
	want := regexp.MustCompile(`b` + name + `b`)
	message, err := CustomHello(name)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`CustomHello("") = %q, %v, want "", error`, message, err)
	}

}
